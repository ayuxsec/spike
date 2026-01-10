package cli

import (
	"testing"

	"github.com/ayuxsec/spike/pkg/config"
)

func TestNucleiInferTargetType(t *testing.T) {
	scanType := NucleiGenericScanType
	t.Log(scanType.inferNucleiTargetType())
	t.Log(scanType.String())
}

func TestRunNucleiGeneric(t *testing.T) {
	cfg := config.DefaultConfig().Tools.Nuclei

	targets := []string{"https://example.com"}

	// Just test that it doesn't error on argument assembly.
	_, err := RunNuclei(targets, &cfg, NucleiGenericScanType)
	if err != nil {
		t.Fatalf("RunNuclei (Generic) returned error: %v", err)
	}
}

func TestRunNucleiDast(t *testing.T) {
	cfg := config.DefaultConfig().Tools.Nuclei

	targets := []string{"https://example.com/?q=test"}

	_, err := RunNuclei(targets, &cfg, NucleiDastScanType)
	if err != nil {
		t.Fatalf("RunNuclei (DAST) returned error: %v", err)
	}
}
