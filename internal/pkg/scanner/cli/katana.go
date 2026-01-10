package cli

import (
	"strconv"

	"github.com/ayuxsec/spike/pkg/config"
	"github.com/ayuxsec/spike/pkg/logger"
)

// katana actively crawls the given targets via interacting with web pages
func RunKatana(targets []string, k *config.KatanaConfig) ([]string, error) {
	logger.Infof("Running katana on %d http server", len(targets))
	if !k.Enabled {
		logger.Warn("Katana is disabled in the config, skipping Katana scan")
		return nil, nil
	}
	katanaArgs := []string{"-jc", "-kf", "all", "-fx", "-xhr", "-jsl", "-aff",
		"-c", strconv.Itoa(k.Threads),
		"-p", strconv.Itoa(k.ParallelismThreads),
		"-d", strconv.Itoa(k.CrawlDepth),
		"-ct", k.MaxCrawlTime,
	}

	if k.Headless {
		katanaArgs = append(katanaArgs, "-headless")
	}
	if k.NoSandbox {
		katanaArgs = append(katanaArgs, "-no-sandbox")
	}
	return RunCommandWithStdinInput("katana", katanaArgs, targets, k.CmdTimeout)
}
