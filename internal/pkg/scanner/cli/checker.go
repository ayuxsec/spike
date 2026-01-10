package cli

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/ayuxsec/spike/pkg/logger"
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
	var isAnyMissing bool
	for _, tool := range requiredToolsList {
		isInstalled, err := ChkCmdInstalled(tool)
		if err != nil {
			logger.Warnf("error checking tool %s installation: %v", tool, err)
			continue
		}
		if !isInstalled {
			isAnyMissing = true
			logger.Warnf("tool '%s' not found in PATH related scan steps will be skipped", tool)
		}
	}
	if isAnyMissing {
		if !askConfirm("Some required tools are missing. Do you still want to continue?") {
			return errors.New("user aborted due to missing required tools")
		}
	}
	return nil
}

func askConfirm(prompt string) bool {
	logger.WarnNoNL(prompt + " [y/N]: ")
	var resp string
	_, _ = fmt.Scanln(&resp)
	resp = strings.ToLower(strings.TrimSpace(resp))
	return resp == "y" || resp == "yes"
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
