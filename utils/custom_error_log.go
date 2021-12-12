package utils

import (
	"errors"
	"fmt"
	"os/exec"
)

func CustomErrorLog(err error) error {
	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		return fmt.Errorf(string(exitError.Stderr))
	} else {
		return fmt.Errorf("%+v", err)
	}
}
