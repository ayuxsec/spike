package cli

import (
	"spike/pkg/config"
	"spike/pkg/logger"
	"strconv"
	"strings"
)

func RunHTTPX(targets []string, cfg *config.HTTPXConfig) ([]string, error) {
	logger.Infof("Running httpx on %d %s", len(targets), inferHttpxTargetType(targets[0]))

	args := []string{"-threads", strconv.Itoa(cfg.Threads)}
	if cfg.ScreenShot {
		args = append(args, "-ss")
	}

	return RunCommandWithStdinInput("httpx", args, targets)
}

func inferHttpxTargetType(url string) string {
	switch {
	case strings.HasPrefix(url, "http://"),
		strings.HasPrefix(url, "https://"):
		return "http servers"
	default:
		return "(sub)domains"
	}
}
