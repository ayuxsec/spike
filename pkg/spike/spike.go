// Package spike provides the main Spike application structure and functionality.
// It attaches the scanner and reporter components to run the scanning process.
package spike

import (
	"fmt"

	"github.com/ayuxsec/spike/internal/pkg/reporter"
	"github.com/ayuxsec/spike/internal/pkg/scanner/core"
	"github.com/ayuxsec/spike/internal/pkg/shell"
	"github.com/ayuxsec/spike/pkg/config"
)

type Scanner struct {
	InputDomains []string
	DBPath       string
	Config       *config.Config
	ScanMode     string // domains to scan (db or input)
}

func NewScanner(opts ...Option) (*Scanner, error) {
	s := &Scanner{}
	for _, opt := range opts {
		opt(s)
	}
	return s, nil
}

func (s *Scanner) Run() error {
	internalScanner, err := s.wire()
	if err != nil {
		return fmt.Errorf("failed to wire package scanner to internal scanner: %w", err)
	}
	if s.Config.Reporter.Telegram.Enabled {
		internalScanner.EventHandler = reporter.NewTelegramNotifier(&s.Config.Reporter)
	}

	// run scanning
	scanErr := internalScanner.ScanDomains()

	// priority: close the scanner even if scanning failed
	closeErr := internalScanner.Close()
	if closeErr != nil {
		return fmt.Errorf("%w: %v", ErrClosingScanner, closeErr)
	}

	if scanErr != nil {
		return fmt.Errorf("%w: %v", ErrFailedScanner, scanErr)
	}

	return nil
}

func NewREPLShell(dbPath string) error {
	return shell.NewREPLShell(dbPath)
}

// wire bulds the internal core.Scanner and wires reporters.
func (s *Scanner) wire() (internalScanner *core.Scanner, err error) {
	scanner, err := core.NewScanner(
		&s.Config.Tools,
		&s.Config.Reporter,
		s.InputDomains,
		s.DBPath,
		mapStrScanMode(s.ScanMode),
	)
	if err != nil {
		return nil, err
	}
	return scanner, nil
}

func mapStrScanMode(mode string) core.ScanMode {
	switch mode {
	case "db":
		return core.DbScanMode
	case "input":
		return core.ManualScanMode
	default:
		return core.ManualScanMode
	}
}
