package cli

import "spike/pkg/logger"

// filters URLs with uro
func RunUro(urls []string) ([]string, error) {
	logger.Infof("Running uro on %d Urls", len(urls))
	return RunCommandWithStdinInput("uro", nil, urls)
}
