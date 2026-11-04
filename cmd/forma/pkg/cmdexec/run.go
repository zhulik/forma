package cmdexec

import (
	"context"
	"os"
	"os/exec"
)

// Run runs name with args under ctx, forwarding stdout and stderr in real time.
// Cancelling ctx kills the child (via exec.CommandContext).
func Run(ctx context.Context, dir, name string, args ...string) error {
	cmd := command(ctx, dir, name, args...)
	return cmd.Run()
}

// Process is a started child with forwarded stdio.
type Process struct {
	cmd *exec.Cmd
}

// Start begins a long-running command without waiting. Cancelling ctx kills the child.
// Call Wait to block until exit.
func Start(ctx context.Context, dir, name string, args ...string) (*Process, error) {
	cmd := command(ctx, dir, name, args...)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return &Process{cmd: cmd}, nil
}

// Wait blocks until the process exits.
func (p *Process) Wait() error {
	return p.cmd.Wait()
}

func command(ctx context.Context, dir, name string, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}
