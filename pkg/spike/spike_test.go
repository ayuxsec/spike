package spike

import (
	"errors"
	"testing"

	"github.com/ayuxsec/spike/pkg/config"
)

func TestRunWithOptions(t *testing.T) {
	cfg := config.DefaultConfig()
	domains := []string{"example.com"}
	scanner, _ := NewScanner(
		OptionWithInputDomains(domains),
		OptionWithDBPath("/tmp/test.db"),
		OptionWithConfig(cfg),
		OptionWithScanMode("input"),
	)
	err := scanner.Run()
	if errors.Is(err, ErrFailedScanner) {
		t.Logf("scanner failed: %v", err)
	} else {
		t.Logf("error closing scanner: %v", err)
	}
}
