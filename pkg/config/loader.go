package config

import (
	"fmt"
	"os"

	"github.com/ayuxsec/spike/pkg/logger"

	"gopkg.in/yaml.v3"
)

func LoadCfg(cfgPath string) (Config, error) {
	logger.Infof("loading config from path: %s", cfgPath)
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config: %v", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to parse config: %v", err)
	}
	return cfg, nil
}

func CreateDefaultCfg(cfgPath string) error {
	data, err := yaml.Marshal(DefaultConfig())
	if err != nil {
		return fmt.Errorf("failed to marshal default config: %v", err)
	}
	if err := os.WriteFile(cfgPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config to %s: %v", cfgPath, err)
	}
	logger.Successf("created %s", cfgPath)
	return nil
}
