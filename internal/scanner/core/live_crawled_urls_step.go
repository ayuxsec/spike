package core

import (
	"spike/internal/scanner/cli"
	"spike/pkg/logger"
)

type LiveCrawledUrlsStep struct{}

func (LiveCrawledUrlsStep) Name() string { return "live_crawled_urls" }

func (LiveCrawledUrlsStep) Enabled(s *Scanner) bool {
	return true // always on
}

func (LiveCrawledUrlsStep) RequiredTools() []string {
	return []string{"uro", "httpx"}
}

func (LiveCrawledUrlsStep) Run(s *Scanner, urls []string) ([]string, error) {

	// 1) Run uro filter
	uroOut, err := cli.RunUro(urls)
	if err != nil {
		return nil, err
	}
	logger.Infof("Recieved %d results from uro, will be passed to httpx to filter dead urls.", len(uroOut))

	// 2) Run httpx to verify live URLs
	live, err := cli.RunHTTPX(uroOut, &s.ToolsCfg.HTTPX)
	if err != nil {
		return nil, err
	}
	logger.Infof("Recieved %d results from httpx.", len(live))

	return live, nil
}

func (LiveCrawledUrlsStep) Store(s *Scanner, out []string) error {
	return s.toolsRepo.LiveCrawledUrls.BulkInsert(s.currentDomain.Id, out)
}

func (LiveCrawledUrlsStep) Fetch(s *Scanner) ([]string, error) {
	return s.toolsRepo.LiveCrawledUrls.Fetch(s.currentDomain.Id)
}
