package cli

import (
	"errors"
	"fmt"
	"os/exec"
	"spike/pkg/logger"
)

// requiredToolsList defines the list of essential external tools needed for scanning.
var requiredToolsList = []string{
	"subfinder",
	"gau",
	"httpx",
	"cachex",
	"katana",
	"nuclei",
	"uro",
}

func WarnIfToolsMissing() error {
	for _, tool := range requiredToolsList {
		isInstalled, err := ChkCmdInstalled(tool)
		if err != nil {
			logger.Warnf("error checking tool %s installation: %v", tool, err)
			continue
		}
		if !isInstalled {
			logger.Warnf("tool '%s' not found in PATH related scan steps will be skipped", tool)
		}
	}
	return nil
}

func ChkCmdInstalled(cmdName string) (bool, error) {
	_, err := exec.LookPath(cmdName)
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return false, nil
		} else {
			return false, fmt.Errorf("error checking command %s installation: %w", cmdName, err)
		}
	}
	return true, nil
}
