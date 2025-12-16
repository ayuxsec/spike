package config

import (
	"fmt"
	"os"
	"path/filepath"

	"spike/pkg/logger"

	"gopkg.in/yaml.v3"
)

// Cfg holds the loaded configuration.
// It is populated by LoadOrCreateConfig and should not be accessed before calling it.
var Cfg *Config

// LoadOrCreateConfig loads the configuration from the given file path.
//
// Behavior:
//   - If the file exists, it is read and unmarshaled into Cfg.
//   - If the file does NOT exist, a new file is created using DefaultConfig().
//   - If forceCreate is true, the existing file (if any) is overwritten with default values.
//
// Returns:
//   - An error if the file cannot be read, created, or parsed.
//
// Notes:
//   - The function ensures that parent directories for the config file exist.
//   - Log messages indicate whether a new config file is created or overwritten.
func LoadOrCreateConfig(cfgPath string, forceCreate bool) error {
	doExist, err := doExistCfg(cfgPath)
	if err != nil {
		return fmt.Errorf("failed to check config file existence: %v", err)
	}

	if !doExist || forceCreate {
		if !doExist {
			logger.Infof("Config file %s does not exist, creating with default values", cfgPath)
		} else {
			logger.Infof("Resetting config file %s with default values", cfgPath)
		}

		if err := createCfgWithDefaultValues(cfgPath); err != nil {
			return fmt.Errorf("failed to create config file: %v", err)
		}
	}

	return loadCfg(cfgPath)
}

// doExistCfg checks whether the given config path points to an existing file.
//
// Returns:
//   - (true, nil) if the file exists
//   - (false, nil) if the file does not exist
//   - (false, err) if an unexpected filesystem error occurs
func doExistCfg(cfgPath string) (bool, error) {
	_, err := os.Stat(cfgPath)
	if err != nil && !os.IsNotExist(err) {
		return false, err
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

// loadCfg reads and unmarshals the YAML configuration from cfgPath into Cfg.
//
// Returns an error if the file cannot be read or parsed.
func loadCfg(cfgPath string) error {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("failed to parse config: %v", err)
	}

	Cfg = &cfg
	return nil
}

// createCfgWithDefaultValues creates (or overwrites) a YAML config file at cfgPath
// using values from DefaultConfig(). The function ensures the parent directory exists.
//
// Returns an error if directory creation, YAML marshaling, or file writing fails.
func createCfgWithDefaultValues(cfgPath string) error {
	data, err := yaml.Marshal(DefaultConfig())
	if err != nil {
		return fmt.Errorf("failed to marshal default config: %v", err)
	}

	if err := ensureCfgDirExist(cfgPath); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	if err := os.WriteFile(cfgPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config to %s: %v", cfgPath, err)
	}

	return nil
}

// ensureCfgDirExist creates the directory tree required for cfgPath.
//
// Returns an error if the directory cannot be created.
func ensureCfgDirExist(cfgPath string) error {
	return os.MkdirAll(getCfgDirPath(cfgPath), 0755)
}

// getCfgDirPath returns the directory component of the config file path.
func getCfgDirPath(cfgPath string) string {
	return filepath.Dir(cfgPath)
}
