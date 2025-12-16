package cli

import (
	"spike/pkg/config"
	"spike/pkg/logger"
	"strconv"
)

// gau passively fetches known URLs for the given targets from various sources
func RunGau(domain string, g *config.GauConfig) ([]string, error) {
	logger.Infof("Running gau on %s", domain)
	if !g.Enabled {
		logger.Warn("Gau is disabled in the config, skipping Gau scan")
		return nil, nil
	}
	args := []string{"--threads", strconv.Itoa(g.Threads), "--subs", domain}
	return RunCommand("gau", args)
}
