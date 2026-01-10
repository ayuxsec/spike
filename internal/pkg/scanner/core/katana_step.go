package core

import "github.com/ayuxsec/spike/internal/pkg/scanner/cli"

type KatanaStep struct{}

func (KatanaStep) Name() string { return "katana" }

func (KatanaStep) Enabled(s *Scanner) bool {
	return s.ToolsCfg.Katana.Enabled
}

func (KatanaStep) RequiredTools() []string {
	return []string{"katana"}
}

func (KatanaStep) Run(s *Scanner, httpServers []string) ([]string, error) {
	return cli.RunKatana(httpServers, &s.ToolsCfg.Katana)
}

func (KatanaStep) Store(s *Scanner, out []string) error {
	return s.toolsRepo.Katana.BulkInsert(s.currentDomain.Id, out)
}

func (KatanaStep) Fetch(s *Scanner) ([]string, error) {
	return s.toolsRepo.Katana.Fetch(s.currentDomain.Id)
}
