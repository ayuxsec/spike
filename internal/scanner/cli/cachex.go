package cli

import (
	"spike/pkg/config"
	"spike/pkg/logger"
	"strconv"
)

func RunCachex(targets []string, c *config.CachexConfig) ([]string, error) {
	logger.Infof("Running cachex on %d http severs", len(targets))
	if !c.Enabled {
		logger.Warn("cachex is disabled in the config, skipping cachex scan")
		return nil, nil
	}
	return RunCommandWithStdinInput("cachex", []string{"-t", strconv.Itoa(c.Threads)}, targets)
}
