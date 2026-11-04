package devcmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/zhulik/forma/cmd/forma/internal/formaconfig"
	"github.com/zhulik/forma/cmd/forma/pkg/cmdexec"
)

// Browser build contract used by forma dev (GoScript → Vite lib bundle):
//
//  1. goscript compile --package . --output <web.path>/.goscript
//  2. npm run build:go  (Vite lib → <web.path>/public/app.js)
//
// from the project / web roots. Vite serves public/ at /, and scaffold main.tsx
// registers globalThis.formaHost.setTree before importing /app.js.
const (
	goscriptBin     = "goscript"
	goscriptOutRel  = ".goscript"
	viteNPMScript   = "dev"
	goBundleScript  = "build:go"
	rebuildDebounce = 200 * time.Millisecond
)

// LongRunner starts a long-lived child and waits for exit (Vite).
type LongRunner func(ctx context.Context, dir, name string, args ...string) error

// GoBuilder runs one browser Go→JS build for the public artifact.
type GoBuilder func(ctx context.Context, projectDir, artifactPath string) error

// PathLooker resolves a binary on PATH (defaults to exec.LookPath).
type PathLooker func(file string) (string, error)

// Runner orchestrates Vite + the GoScript rebuild loop (implements pal.Runner).
type Runner struct {
	Dir    string
	Logger *slog.Logger

	// Optional test seams; nil uses production defaults.
	LookPath PathLooker
	RunLong  LongRunner
	BuildGo  GoBuilder
	WatchGo  func(ctx context.Context, roots []string, onChange func()) error
}

// Run starts the browser development loop until ctx is cancelled.
func (r *Runner) Run(ctx context.Context) error {
	target, err := resolveDir(r.Dir)
	if err != nil {
		return err
	}

	project, webRoot, artifact, err := r.prepare(ctx, target)
	if err != nil {
		return err
	}
	_ = project

	runLong := r.RunLong
	if runLong == nil {
		runLong = defaultRunLong
	}
	buildGo := r.BuildGo
	if buildGo == nil {
		buildGo = defaultBuildGo
	}
	watchGo := r.WatchGo
	if watchGo == nil {
		watchGo = watchGoSources
	}

	r.Logger.InfoContext(ctx, "forma dev: building browser Go artifact", "out", artifact)
	if err := buildGo(ctx, target, artifact); err != nil {
		return fmt.Errorf("initial GoScript build: %w", err)
	}
	r.Logger.InfoContext(ctx, "forma dev: initial GoScript build ok")

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		r.Logger.InfoContext(gctx, "forma dev: starting Vite", "web", webRoot, "script", "npm run "+viteNPMScript)
		err := runLong(gctx, webRoot, "npm", "run", viteNPMScript)
		if gctx.Err() != nil {
			return nil
		}
		if err != nil {
			return fmt.Errorf("vite (npm run %s): %w", viteNPMScript, err)
		}
		return errors.New("vite exited unexpectedly")
	})

	g.Go(func() error {
		return runRebuildLoop(gctx, r.Logger, target, artifact, buildGo, watchGo, rebuildDebounce)
	})

	err = g.Wait()
	if errors.Is(err, context.Canceled) || ctx.Err() != nil {
		r.Logger.InfoContext(context.Background(), "forma dev: stopped")
		return nil
	}
	return err
}

func (r *Runner) prepare(ctx context.Context, target string) (formaconfig.Project, string, string, error) {
	lookPath := r.LookPath
	if lookPath == nil {
		lookPath = defaultLookPath
	}

	project, webRoot, err := CheckPrerequisites(target, lookPath)
	if err != nil {
		return formaconfig.Project{}, "", "", err
	}

	artifact := project.BrowserAppJSPath(target)
	if err := os.MkdirAll(filepath.Dir(artifact), 0o755); err != nil {
		return formaconfig.Project{}, "", "", fmt.Errorf("create browser artifact dir: %w", err)
	}

	r.Logger.InfoContext(ctx, "forma dev: project ready",
		"dir", target,
		"web", webRoot,
		"artifact", artifact,
	)
	return project, webRoot, artifact, nil
}

func resolveDir(directory string) (string, error) {
	if directory == "" {
		return os.Getwd()
	}
	return filepath.Abs(directory)
}

func defaultRunLong(ctx context.Context, dir, name string, args ...string) error {
	proc, err := cmdexec.Start(ctx, dir, name, args...)
	if err != nil {
		return err
	}
	return proc.Wait()
}

func defaultBuildGo(ctx context.Context, projectDir, artifactPath string) error {
	webRoot := filepath.Dir(filepath.Dir(artifactPath)) // …/web/public/app.js → …/web
	outDir := filepath.Join(webRoot, goscriptOutRel)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}
	if err := cmdexec.Run(ctx, projectDir, goscriptBin, "compile", "--package", ".", "--output", outDir); err != nil {
		return fmt.Errorf("goscript compile: %w", err)
	}
	if err := cmdexec.Run(ctx, webRoot, "npm", "run", goBundleScript); err != nil {
		return fmt.Errorf("npm run %s: %w", goBundleScript, err)
	}
	return nil
}
