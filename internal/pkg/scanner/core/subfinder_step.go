package core

import "github.com/ayuxsec/spike/internal/pkg/scanner/cli"

type SubfinderStep struct{}

func (SubfinderStep) Name() string { return "subfinder" }

func (SubfinderStep) Enabled(s *Scanner) bool {
	return s.ToolsCfg.Subfinder.Enabled
}

func (SubfinderStep) RequiredTools() []string {
	return []string{"subfinder"}
}

func (SubfinderStep) Run(s *Scanner, rootDomain []string) ([]string, error) {
	return cli.RunSubfinder(rootDomain[0], &s.ToolsCfg.Subfinder)
}

func (SubfinderStep) Store(s *Scanner, out []string) error {
	return s.toolsRepo.Subfinder.BulkInsert(s.currentDomain.Id, out)
}

func (SubfinderStep) Fetch(s *Scanner) ([]string, error) {
	return s.toolsRepo.Subfinder.Fetch(s.currentDomain.Id)
}
