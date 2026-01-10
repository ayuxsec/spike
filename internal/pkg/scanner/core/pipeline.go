package core

import "strings"

type PipelineStep struct {
	Step      ToolStep
	InputFrom []string
}

func (s *Scanner) selectPipeline() []PipelineStep {
	if strings.HasPrefix(s.currentDomain.Name, "*.") { // domain starting with *. indicates wildcard
		return WildCardDomainPipeline()
	}
	return SingleDomainPipeline()
}

func WildCardDomainPipeline() []PipelineStep {
	return []PipelineStep{
		{Step: SubfinderStep{}, InputFrom: []string{"__root_domain__"}},
		{Step: HttpxStep{}, InputFrom: []string{"subfinder"}},
		{Step: GauStep{}, InputFrom: []string{"__root_domain__"}},
		{Step: CachexStep{}, InputFrom: []string{"httpx"}},
		{Step: KatanaStep{}, InputFrom: []string{"httpx"}},
		{Step: UroStep{}, InputFrom: []string{"katana", "gau"}},
		{Step: NucleiGenericStep{}, InputFrom: []string{"httpx"}},
		{Step: NucleiDastStep{}, InputFrom: []string{"uro"}},
	}
}

func SingleDomainPipeline() []PipelineStep {
	return []PipelineStep{
		{Step: HttpxStep{}, InputFrom: []string{"__root_domain__"}},
		{Step: GauStep{}, InputFrom: []string{"__root_domain__"}},
		{Step: CachexStep{}, InputFrom: []string{"httpx"}},
		{Step: KatanaStep{}, InputFrom: []string{"httpx"}},
		{Step: UroStep{}, InputFrom: []string{"katana", "gau"}},
		{Step: NucleiGenericStep{}, InputFrom: []string{"httpx"}},
		{Step: NucleiDastStep{}, InputFrom: []string{"uro"}},
	}
}
