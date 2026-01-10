package core

import "github.com/ayuxsec/spike/internal/pkg/scanner/cli"

type NucleiGenericStep struct{}

func (NucleiGenericStep) Name() string { return "nuclei_generic" }

func (NucleiGenericStep) Enabled(s *Scanner) bool {
	return s.ToolsCfg.Nuclei.TemplateSettings.Generic
}

func (NucleiGenericStep) RequiredTools() []string {
	return []string{"nuclei"}
}

func (NucleiGenericStep) Run(s *Scanner, input []string) ([]string, error) {
	return cli.RunNuclei(input, &s.ToolsCfg.Nuclei, cli.NucleiGenericScanType)
}

func (NucleiGenericStep) Store(s *Scanner, out []string) error {
	return s.toolsRepo.Nuclei.BulkInsert(s.currentDomain.Id, out)
}

func (NucleiGenericStep) Fetch(s *Scanner) ([]string, error) {
	return s.toolsRepo.Nuclei.Fetch(s.currentDomain.Id)
}
