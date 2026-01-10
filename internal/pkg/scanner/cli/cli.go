package cli

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"

	"github.com/ayuxsec/spike/pkg/logger"
)

// RunCommand runs a command and returns stdout lines with duplicates and empty lines removed.
func RunCommand(
	cmdName string,
	args []string,
	timeout int,
) ([]string, error) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(timeout)*time.Second,
	)
	defer cancel()

	var stdout bytes.Buffer

	cmd := exec.CommandContext(ctx, cmdName, args...)
	cmd.Stdout = &stdout
	//cmd.Stderr = &stderr

	logger.Debugf("Running command: %s", cmd.String())

	err := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("command timed out after %ds: %w", timeout, ErrCtxTimedOut)
	}

	if err != nil {
		return nil, fmt.Errorf("command failed: %v", err)
	}

	lines := LinesToSlice(stdout.String())
	return RemoveDuplicatesAndEmptyStrings(lines), nil
}

// RunCommandWithStdinInput runs a command with StdinInput piped to stdin,
// and returns stdout lines with duplicates and empty lines removed.
func RunCommandWithStdinInput(
	cmdName string,
	args []string,
	stdinInput []string,
	timeout int, // timeout in seconds
) ([]string, error) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(timeout)*time.Second,
	)
	defer cancel()

	var stdout bytes.Buffer

	cmd := exec.CommandContext(ctx, cmdName, args...)
	cmd.Stdout = &stdout
	//cmd.Stderr = &stderr: todo handle stderr

	logger.Debugf("Running command: %s", cmd.String())

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get stdin: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start command: %v", err)
	}

	// Write stdin in a goroutine so cmd.Wait() can't deadlock
	go func() {
		defer stdin.Close()
		input := strings.Join(stdinInput, "\n") + "\n"
		_, _ = io.WriteString(stdin, input)
	}()

	err = cmd.Wait()

	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("command timed out after %ds: %w", timeout, ErrCtxTimedOut)
	}

	if err != nil {
		return nil, fmt.Errorf("command failed: %v", err)
	}

	lines := LinesToSlice(stdout.String())
	return RemoveDuplicatesAndEmptyStrings(lines), nil
}
