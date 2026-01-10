package spike

import "github.com/ayuxsec/spike/pkg/config"

type Option func(s *Scanner)

// OptionWithInputDomains sets the input domains.
func OptionWithInputDomains(inputDomains []string) Option {
	return func(s *Scanner) {
		s.InputDomains = inputDomains
	}
}

// OptionWithDBPath sets the database path.
func OptionWithDBPath(dbPath string) Option {
	return func(s *Scanner) {
		s.DBPath = dbPath
	}
}

// OptionWithConfig sets the config.
func OptionWithConfig(cfg *config.Config) Option {
	return func(s *Scanner) {
		s.Config = cfg
	}
}

// OptionWithScanMode sets the scan mode.
func OptionWithScanMode(mode string) Option {
	return func(s *Scanner) {
		s.ScanMode = mode
	}
}
