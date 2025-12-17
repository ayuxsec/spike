// Package spike provides the main Spike application structure and functionality.
// It attaches the scanner and reporter components to run the scanning process.
package spike

import (
	"fmt"
	"spike/internal/reporter"
	"spike/internal/scanner/core"
	"spike/pkg/config"
)

type Spike struct {
	Scanner     *core.Scanner
	ReporterCfg *config.ReporterConfig
}

func New(
	toolsCfg *config.ToolsConfig,
	reporterCfg *config.ReporterConfig,
	inputDomains []string,
	dbPath string,
	scanMode core.ScanMode,
) (*Spike, error) {

	scanner, err := core.NewScanner(
		toolsCfg,
		reporterCfg,
		inputDomains,
		dbPath,
		scanMode,
	)
	if err != nil {
		return nil, err
	}

	return &Spike{
		Scanner:     scanner,
		ReporterCfg: reporterCfg,
	}, nil
}

func (s *Spike) Run() error {

	if s.ReporterCfg.Telegram.Enabled {
		s.Scanner.EventHandler = reporter.NewTelegramNotifier(s.ReporterCfg)
	}

	// run scanning
	scanErr := s.Scanner.ScanDomains()

	// priority: close the scanner even if scanning failed
	closeErr := s.Scanner.Close()
	if closeErr != nil {
		return fmt.Errorf("failed to close scanner: %v", closeErr)
	}

	if scanErr != nil {
		return scanErr
	}

	return nil
}
