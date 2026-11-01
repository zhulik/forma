package cmdexec

import (
	"context"
	"os"
	"os/exec"
)

// Run runs name with args under ctx, forwarding stdout and stderr in real time.
func Run(ctx context.Context, dir, name string, args ...string) error {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
