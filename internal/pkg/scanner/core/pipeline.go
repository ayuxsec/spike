package core

import (
	"strings"
)

type PipelineStep struct {
	Step      ToolStep
	InputFrom []string
}

func (s *Scanner) selectPipeline() []PipelineStep {
	var isWildCard bool
	// Trim the "*." prefix and restore the base domain; it is only needed for pipeline selection.
	s.currentDomain.Name, isWildCard = strings.CutPrefix(s.currentDomain.Name, "*.") // *. is a magic string; domain starting with *. indicates wildcard
	if isWildCard {
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
