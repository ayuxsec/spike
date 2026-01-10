package shell

import "testing"

func TestNewREPLShell(t *testing.T) {
	if err := NewREPLShell("spike.db"); err != nil {
		t.Log(err)
	}
}
