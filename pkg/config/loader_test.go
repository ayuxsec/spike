package config

import (
	"testing"
)

func TestLoadOrCreateConfig(t *testing.T) {
	// Test with a valid config file
	err := LoadOrCreateConfig("example_config.yaml", false)
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}
	if Cfg == nil {
		t.Fatal("LoadConfig() did not load config")
	}
	t.Logf("LoadConfig() loaded config: %+v", Cfg)
}
