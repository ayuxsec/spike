package cli

import (
	"errors"
	"testing"
)

func TestRunCommand(t *testing.T) {
	var args []string
	for range 100 {
		args = append(args, "test\n")
	}
	output, err := RunCommand("echo", args, 1)
	if err != nil {
		if errors.Is(err, ErrCtxTimedOut) {
			t.Log("command timed out, partial output will be given")
		} else if errors.Is(err, ErrSignalInterrupt) {
			t.Log("user pressed ctrl+c, partial output will be given")
		}
	}
	t.Log(output)
}
