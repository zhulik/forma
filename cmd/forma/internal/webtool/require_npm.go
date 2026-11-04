package webtool

import (
	"errors"
	"os/exec"
)

func requireNpm() error {
	if _, err := exec.LookPath("npm"); err != nil {
		return errors.New("npm not found on PATH; install Node.js to sync the web tree")
	}

	return nil
}
