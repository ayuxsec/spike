package cli

import (
	"spike/pkg/config"
	"spike/pkg/logger"
	"strconv"
	"strings"
)

type NucleiScanType int

const (
	NucleiDastScanType NucleiScanType = iota
	NucleiGenericScanType
)

func (nst NucleiScanType) String() string {
	switch nst {
	case NucleiDastScanType:
		return "DAST"
	case NucleiGenericScanType:
		return "Generic"
	}
	return "Unknown"
}

func RunNuclei(targets []string, cfg *config.NucleiConfig, scanType NucleiScanType) ([]string, error) {
	var args []string
	var targetType string
	switch scanType {
	case NucleiDastScanType:
		args = append(args, "-dast")
		targetType = "crawled urls" // expected target to be passed when run in DAST mode
	case NucleiGenericScanType:
		targetType = "http servers" // expected target to be passed when run in Generic mode
	}

	logger.Infof("Running Nuclei %s scan on %d %s", scanType.String(), len(targets), targetType)

	include := strings.TrimRight(strings.Join(cfg.TemplatePaths.Generic.Include, ","), ",")
	exclude := strings.TrimRight(strings.Join(cfg.TemplatePaths.Generic.Exclude, ","), ",")

	cliArgs := []string{
		"-c", strconv.Itoa(cfg.Threads),
		"-it", include,
		"-et", exclude,
	}

	args = append(args, cliArgs...)

	if cfg.TemplateSettings.Headless {
		args = append(args, "-headless")
	}

	return RunCommandWithStdinInput("nuclei", args, targets)
}
