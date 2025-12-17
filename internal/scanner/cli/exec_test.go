package cli

import "testing"

func TestRunCommandWithStdinInput(t *testing.T) {
	cmdName := "cat"
	stdinInput := []string{"test", "example"}
	out, err := RunCommandWithStdinInput(cmdName, nil, stdinInput, 10)
	if err != nil {
		t.Fatalf("TestRunCommandWithStdinInput failed: %v", err)
	}
	t.Logf("TestRunCommandWithStdinInput output: %v", out)
}
