package core

import (
	"github.com/ayuxsec/spike/internal/pkg/scanner/cli"
	"github.com/ayuxsec/spike/pkg/logger"
)

type UroStep struct{}

func (UroStep) Name() string { return "uro" }

func (UroStep) Enabled(s *Scanner) bool {
	return true // always on
}

func (UroStep) RequiredTools() []string {
	return []string{"uro", "httpx"}
}

func (UroStep) Run(s *Scanner, urls []string) ([]string, error) {

	// 1) Run uro filter
	uroOut, err := cli.RunUro(urls)
	if err != nil {
		return nil, err
	}
	logger.Infof("Recieved %d results from uro, will be passed to httpx to filter dead urls.", len(uroOut))

	// 2) Run httpx to verify live URLs
	live, err := cli.RunHTTPX(uroOut, &s.ToolsCfg.HTTPX, false)
	if err != nil {
		return nil, err
	}
	logger.Infof("Recieved %d results from httpx.", len(live))

	return live, nil
}

func (UroStep) Store(s *Scanner, out []string) error {
	return s.toolsRepo.Uro.BulkInsert(s.currentDomain.Id, out)
}

func (UroStep) Fetch(s *Scanner) ([]string, error) {
	return s.toolsRepo.Uro.Fetch(s.currentDomain.Id)
}
