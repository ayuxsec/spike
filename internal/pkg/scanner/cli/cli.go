package cli

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"time"

	"github.com/ayuxsec/spike/pkg/logger"
)

// RunCommand runs a command and returns stdout lines with duplicates and empty lines removed.
func RunCommand(cmdName string, args []string, timeout int) ([]string, error) {
	c := make(chan os.Signal, 1) // channel to track os.Interrupt (`Ctrl + C`)
	signal.Notify(c, os.Interrupt)
	defer signal.Stop(c)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(timeout)*time.Second,
	)
	defer cancel()
	cmd := exec.CommandContext(ctx, cmdName, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %v", err)
	}

	logger.Debugf("Running command: %s", cmd.String())

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start command: %v", err)
	}

	var buf bytes.Buffer
	readDone := make(chan error, 1)

	go func() {
		_, err := io.Copy(&buf, stdout)
		readDone <- err
	}()

	select {
	case <-c:
		_ = cmd.Process.Kill() // todo: handle error
		<-readDone
		lines := LinesToSlice(buf.String())
		return RemoveDuplicatesAndEmptyStrings(lines),
			fmt.Errorf("command interrupted by user: %w", ErrSignalInterrupt)

	case <-ctx.Done():
		_ = cmd.Process.Kill()
		<-readDone
		lines := LinesToSlice(buf.String())
		return RemoveDuplicatesAndEmptyStrings(lines),
			fmt.Errorf("command timed out after %ds: %w", timeout, ErrCtxTimedOut)

	case err := <-readDone:
		if err != nil {
			return nil, fmt.Errorf("failed to read stdout: %v", err)
		}
	}

	if err := cmd.Wait(); err != nil {
		logger.Debugf("command exited with error: %v", err)
	}

	lines := LinesToSlice(buf.String())
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

	c := make(chan os.Signal, 1) // channel to track os.Interrupt (`Ctrl + C`)
	signal.Notify(c, os.Interrupt)
	defer signal.Stop(c)

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

	waitDone := make(chan error, 1)
	go func() {
		waitDone <- cmd.Wait()
	}()

	select {
	case <-c:
		_ = cmd.Process.Kill()
		<-waitDone
		lines := LinesToSlice(stdout.String())
		return RemoveDuplicatesAndEmptyStrings(lines),
			fmt.Errorf("command interrupted by user: %w", ErrSignalInterrupt)

	case <-ctx.Done():
		_ = cmd.Process.Kill()
		<-waitDone
		lines := LinesToSlice(stdout.String())
		return RemoveDuplicatesAndEmptyStrings(lines),
			fmt.Errorf("command timed out after %ds: %w", timeout, ErrCtxTimedOut)

	case err := <-waitDone:
		if err != nil {
			return nil, fmt.Errorf("command failed: %v", err)
		}
	}

	lines := LinesToSlice(stdout.String())
	return RemoveDuplicatesAndEmptyStrings(lines), nil
}
