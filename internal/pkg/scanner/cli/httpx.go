package cli

import (
	"strconv"
	"strings"

	"github.com/ayuxsec/spike/pkg/config"
	"github.com/ayuxsec/spike/pkg/logger"
)

func RunHTTPX(targets []string, cfg *config.HTTPXConfig, scanPorts bool) ([]string, error) {
	if len(targets) == 0 {
		return nil, nil
	}
	logger.Infof("Running httpx on %d %s", len(targets), inferHttpxTargetType(targets[0]))

	args := []string{"-threads", strconv.Itoa(cfg.Threads)}
	if cfg.ScreenShot {
		args = append(args, "-ss")
	}

	if scanPorts {
		if cfg.TargetPorts.Http != "" {
			args = append(args, "-ports", "http:"+cfg.TargetPorts.Http)
		}
		if cfg.TargetPorts.Https != "" {
			args = append(args, "-ports", "https:"+cfg.TargetPorts.Https)
		}
	}

	return RunCommandWithStdinInput("httpx", args, targets, cfg.CmdTimeout)
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
