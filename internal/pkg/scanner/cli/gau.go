package cli

import (
	"strconv"
	"strings"

	"github.com/ayuxsec/spike/pkg/config"
	"github.com/ayuxsec/spike/pkg/logger"
)

// gau passively fetches known URLs for the given targets from various sources
func RunGau(domain string, g *config.GauConfig) ([]string, error) {
	logger.Infof("Running gau on %s", domain)
	if !g.Enabled {
		logger.Warn("Gau is disabled in the config, skipping Gau scan")
		return nil, nil
	}
	args := []string{"--threads", strconv.Itoa(g.Threads)}
	if strings.HasPrefix(domain, "*.") {
		args = append(args, "--subs", domain)
	} else {
		args = append(args, domain)
	}
	return RunCommand("gau", args, g.CmdTimeout)
}
