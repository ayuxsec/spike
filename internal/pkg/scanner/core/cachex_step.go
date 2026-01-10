package core

import "github.com/ayuxsec/spike/internal/pkg/scanner/cli"

type CachexStep struct{}

func (CachexStep) Name() string { return "cachex" }

func (CachexStep) Enabled(s *Scanner) bool {
	return s.ToolsCfg.Cachex.Enabled
}

func (CachexStep) RequiredTools() []string {
	return []string{"cachex"}
}

func (CachexStep) Run(s *Scanner, httpServers []string) ([]string, error) {
	return cli.RunCachex(httpServers, &s.ToolsCfg.Cachex)
}

func (CachexStep) Store(s *Scanner, out []string) error {
	return s.toolsRepo.Cachex.BulkInsert(s.currentDomain.Id, out)
}

func (CachexStep) Fetch(s *Scanner) ([]string, error) {
	return s.toolsRepo.Cachex.Fetch(s.currentDomain.Id)
}
