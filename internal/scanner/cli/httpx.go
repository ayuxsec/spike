package cli

import (
	"spike/pkg/config"
	"spike/pkg/logger"
	"strconv"
	"strings"
)

func RunHTTPX(targets []string, cfg *config.HTTPXConfig) ([]string, error) {
	var targetType string
	switch {
	case strings.HasPrefix(targets[0], "http://"),
		strings.HasPrefix(targets[0], "https://"):
		targetType = "http servers"
	default:
		targetType = "(sub)domains"
	}

	logger.Infof("Running httpx on %d %s", len(targets), targetType)

	args := []string{"-threads", strconv.Itoa(cfg.Threads)}
	if cfg.ScreenShot {
		args = append(args, "-ss")
	}

	return RunCommandWithStdinInput("httpx", args, targets)
}
