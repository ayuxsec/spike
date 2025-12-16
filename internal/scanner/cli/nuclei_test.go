package cli

import (
	"spike/pkg/config"
	"testing"
)

func TestRunNucleiGeneric(t *testing.T) {
	cfg := config.DefaultConfig().ToolsConfig.Nuclei

	targets := []string{"https://example.com"}

	// Just test that it doesn't error on argument assembly.
	_, err := RunNuclei(targets, &cfg, NucleiGenericScanType)
	if err != nil {
		t.Fatalf("RunNuclei (Generic) returned error: %v", err)
	}
}

func TestRunNucleiDast(t *testing.T) {
	cfg := config.DefaultConfig().ToolsConfig.Nuclei

	targets := []string{"https://example.com/?q=test"}

	_, err := RunNuclei(targets, &cfg, NucleiDastScanType)
	if err != nil {
		t.Fatalf("RunNuclei (DAST) returned error: %v", err)
	}
}
