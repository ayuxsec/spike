package cli

import (
	"strconv"
	"strings"

	"github.com/ayuxsec/spike/pkg/config"
	"github.com/ayuxsec/spike/pkg/logger"
)

type NucleiScanType int

const (
	_ NucleiScanType = iota
	NucleiDastScanType
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

func (nst NucleiScanType) inferNucleiTargetType() string {
	switch nst {
	case NucleiDastScanType:
		return "crawled urls" // expected target to be passed when run in DAST mode
	case NucleiGenericScanType:
		return "http servers" // expected target to be passed when run in Generic mode
	}
	return "Unknown"
}

func RunNuclei(targets []string, cfg *config.NucleiConfig, scanType NucleiScanType) ([]string, error) {
	var args []string
	if scanType == NucleiDastScanType {
		args = append(args, "-dast")
	}

	logger.Infof("Running Nuclei %s Scan on %d %s", scanType.String(), len(targets),
		scanType.inferNucleiTargetType())

	include := strings.TrimRight(strings.Join(cfg.TemplatePaths.Generic.Include, ","), ",")
	exclude := strings.TrimRight(strings.Join(cfg.TemplatePaths.Generic.Exclude, ","), ",")
	excludeSeverity := strings.TrimRight(strings.Join(cfg.TemplatePaths.Generic.ExcludeSeverity, ","), ",")

	cliArgs := []string{
		"-c", strconv.Itoa(cfg.Threads),
		"-it", include,
		"-et", exclude,
		"-es", excludeSeverity,
		"-nc", "-silent",
	}

	args = append(args, cliArgs...)

	if cfg.TemplateSettings.Headless {
		args = append(args, "-headless")
	}

	return RunCommandWithStdinInput("nuclei", args, targets, cfg.CmdTimeout)
}
