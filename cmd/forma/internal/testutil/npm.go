package testutil

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// RequireNpm skips the test when npm cannot run a minimal install (missing Node, broken shim, etc.).
func RequireNpm(t *testing.T) {
	t.Helper()

	if _, err := exec.LookPath("npm"); err != nil {
		t.Skip("npm not on PATH")
	}

	dir := t.TempDir()
	manifest := filepath.Join(dir, "package.json")
	if err := os.WriteFile(manifest, []byte(`{"name":"forma-npm-probe","private":true}`), 0o600); err != nil {
		t.Fatalf("write probe package.json: %v", err)
	}

	cmd := exec.Command("npm", "install", "--ignore-scripts", "--no-audit", "--no-fund")
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		t.Skip("npm install probe failed")
	}
}
