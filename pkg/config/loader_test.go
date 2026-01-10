package config

import (
	"testing"
)

func TestCreateCfg(t *testing.T) {
	if err := CreateDefaultCfg("config_test.yaml"); err != nil {
		t.Log(err)
	}
}

func TestLoadCfg(t *testing.T) {
	if cfg, err := LoadCfg("config_test.yaml"); err != nil {
		t.Log(err)
	} else {
		t.Log(cfg)
	}
}
