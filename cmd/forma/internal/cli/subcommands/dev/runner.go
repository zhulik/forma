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

// P3c GopherJS build contract used by forma dev:
//
//	gopherjs build -o <web.path>/public/app.js .
//
// from the project root. Vite serves public/ at /, and scaffold main.tsx loads
// /app.js only after registering globalThis.formaHost.setTree.
const (
	gopherjsBin     = "gopherjs"
	viteNPMScript   = "dev"
	rebuildDebounce = 200 * time.Millisecond
)

// LongRunner starts a long-lived child and waits for exit (Vite).
type LongRunner func(ctx context.Context, dir, name string, args ...string) error

// GoBuilder runs one GopherJS build for the browser artifact.
type GoBuilder func(ctx context.Context, projectDir, artifactPath string) error

// PathLooker resolves a binary on PATH (defaults to exec.LookPath).
type PathLooker func(file string) (string, error)

// Runner orchestrates Vite + the GopherJS rebuild loop (implements pal.Runner).
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
		return fmt.Errorf("initial GopherJS build: %w", err)
	}
	r.Logger.InfoContext(ctx, "forma dev: initial GopherJS build ok")

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
		return formaconfig.Project{}, "", "", fmt.Errorf("create GopherJS output dir: %w", err)
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
	if err := os.MkdirAll(filepath.Dir(artifactPath), 0o755); err != nil {
		return err
	}
	return cmdexec.Run(ctx, projectDir, gopherjsBin, "build", "-o", artifactPath, ".")
}
