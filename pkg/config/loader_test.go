package config

import (
	"os"
	"testing"
)

func TestLoadOrCreateConfig(t *testing.T) {
	cfgFileName := "example_config.yaml"
	_, err := os.Stat(cfgFileName)
	if err == nil {
		err = os.Remove(cfgFileName)
		if err != nil {
			t.Fatalf("Failed to remove existing config file: %v", err)
		}
	}
	// Test with a valid config file
	err = LoadOrCreateConfig("example_config.yaml", false)
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}
	if Cfg == nil {
		t.Fatal("LoadConfig() did not load config")
	}
	t.Logf("LoadConfig() loaded config: %+v", Cfg)
}
