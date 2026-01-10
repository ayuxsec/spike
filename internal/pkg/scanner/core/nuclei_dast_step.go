package core

import "github.com/ayuxsec/spike/internal/pkg/scanner/cli"

type NucleiDastStep struct{}

func (NucleiDastStep) Name() string { return "nuclei_dast" }

func (NucleiDastStep) Enabled(s *Scanner) bool {
	return s.ToolsCfg.Nuclei.TemplateSettings.Dast
}

func (NucleiDastStep) RequiredTools() []string {
	return []string{"nuclei"}
}

func (NucleiDastStep) Run(s *Scanner, input []string) ([]string, error) {
	return cli.RunNuclei(filterJsEndpoints(input), &s.ToolsCfg.Nuclei, cli.NucleiDastScanType)
}

func (NucleiDastStep) Store(s *Scanner, out []string) error {
	return s.toolsRepo.Nuclei.BulkInsert(s.currentDomain.Id, out)
}

func (NucleiDastStep) Fetch(s *Scanner) ([]string, error) {
	return s.toolsRepo.Nuclei.Fetch(s.currentDomain.Id)
}
