package core

import "github.com/ayuxsec/spike/internal/pkg/scanner/cli"

type GauStep struct{}

func (GauStep) Name() string { return "gau" }

func (GauStep) Enabled(s *Scanner) bool {
	return s.ToolsCfg.Gau.Enabled
}

func (GauStep) RequiredTools() []string {
	return []string{"gau"}
}

func (GauStep) Run(s *Scanner, rootDomain []string) ([]string, error) {
	return cli.RunGau(rootDomain[0], &s.ToolsCfg.Gau)
}

func (GauStep) Store(s *Scanner, out []string) error {
	return s.toolsRepo.Gau.BulkInsert(s.currentDomain.Id, out)
}

func (GauStep) Fetch(s *Scanner) ([]string, error) {
	return s.toolsRepo.Gau.Fetch(s.currentDomain.Id)
}
