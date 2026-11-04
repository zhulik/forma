package devcmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	clierrors "github.com/zhulik/forma/cmd/forma/internal/cli/errors"
	"github.com/zhulik/forma/cmd/forma/internal/formaconfig"
)

var (
	// ErrNodeModulesMissing means web deps were not installed (run forma sync).
	ErrNodeModulesMissing = errors.New("web node_modules missing; run forma sync first")
	// ErrGopherJSMissing means the GopherJS toolchain is not on PATH.
	ErrGopherJSMissing = errors.New("gopherjs not found on PATH; install GopherJS (see doc/user/install.md)")
)

func defaultLookPath(file string) (string, error) {
	return exec.LookPath(file)
}

// CheckPrerequisites loads forma.yaml and fails fast on missing deps/toolchain.
func CheckPrerequisites(projectDir string, lookPath PathLooker) (formaconfig.Project, string, error) {
	if lookPath == nil {
		lookPath = defaultLookPath
	}
	configPath := filepath.Join(projectDir, formaconfig.FileName)
	if _, err := os.Stat(configPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return formaconfig.Project{}, "", fmt.Errorf("%s: %w", formaconfig.FileName, err)
		}
		return formaconfig.Project{}, "", fmt.Errorf("stat %s: %w", formaconfig.FileName, err)
	}

	project, err := formaconfig.Load(projectDir)
	if err != nil {
		return formaconfig.Project{}, "", fmt.Errorf("load %s: %w", formaconfig.FileName, err)
	}

	webRoot := project.WebRoot(projectDir)
	packageJSON := filepath.Join(webRoot, "package.json")
	if _, err := os.Stat(packageJSON); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return formaconfig.Project{}, "", fmt.Errorf("%w at %s", clierrors.ErrWebPackageMissing, packageJSON)
		}
		return formaconfig.Project{}, "", fmt.Errorf("stat web package.json: %w", err)
	}

	nodeModules := filepath.Join(webRoot, "node_modules")
	if st, err := os.Stat(nodeModules); err != nil || !st.IsDir() {
		return formaconfig.Project{}, "", fmt.Errorf("%w under %s", ErrNodeModulesMissing, webRoot)
	}

	if _, err := lookPath(gopherjsBin); err != nil {
		return formaconfig.Project{}, "", ErrGopherJSMissing
	}
	if _, err := lookPath("npm"); err != nil {
		return formaconfig.Project{}, "", errors.New("npm not found on PATH; install Node.js to run forma dev")
	}

	return project, webRoot, nil
}
