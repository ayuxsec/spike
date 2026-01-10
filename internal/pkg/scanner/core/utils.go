package core

import (
	"fmt"
)

func (s *Scanner) checkAllToolsScanned() (bool, error) {
	for _, step := range s.selectPipeline() {
		toolName := step.Step.Name()
		completed, err := s.scanTrackerRepo.IsScanCompleted(s.currentDomain.Id, toolName)
		if err != nil {
			return false, fmt.Errorf("failed to check tool %s completition status from db: %v", toolName, err)
		}
		if !completed {
			return false, nil
		}
	}
	return true, nil
}
