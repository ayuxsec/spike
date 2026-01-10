package core

import "github.com/ayuxsec/spike/internal/pkg/scanner/cli"

type HttpxStep struct{}

func (HttpxStep) Name() string { return "httpx" }

func (HttpxStep) Enabled(s *Scanner) bool { return true }

func (HttpxStep) RequiredTools() []string {
	return []string{"httpx"}
}

func (HttpxStep) Run(s *Scanner, subs []string) ([]string, error) {
	return cli.RunHTTPX(subs, &s.ToolsCfg.HTTPX, true)
}

func (HttpxStep) Store(s *Scanner, out []string) error {
	return s.toolsRepo.Httpx.BulkInsert(s.currentDomain.Id, out)
}

func (HttpxStep) Fetch(s *Scanner) ([]string, error) {
	return s.toolsRepo.Httpx.Fetch(s.currentDomain.Id)
}
