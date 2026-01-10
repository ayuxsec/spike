package core

import (
	"fmt"

	"github.com/ayuxsec/spike/internal/pkg/scanner/cli"
	"github.com/ayuxsec/spike/pkg/logger"
)

func (s *Scanner) ExecStep(step ToolStep, input []string) ([]string, error) {
	domain := s.currentDomain

	if !step.Enabled(s) {
		logger.Infof("Step %s is disabled, skipping...", step.Name())
		return nil, nil
	}

	requiredTools := step.RequiredTools()
	for _, tool := range requiredTools {
		isToolInstalled, err := cli.ChkCmdInstalled(tool)
		if err != nil {
			return nil, fmt.Errorf("failed to check if tool %s is installed: %v", tool, err)
		}
		if !isToolInstalled {
			return nil, fmt.Errorf("tool %s is not installed or not found in PATH, skipping step", tool)
		}
	}

	logger.Infof("Running step: %s", step.Name())

	completed, err := s.scanTrackerRepo.IsScanCompleted(domain.Id, step.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to check scan completiton status %v", err)
	}
	if !completed {
		out, err := step.Run(s, input)
		if err != nil {
			return nil, err
		}

		if len(out) > 0 {
			if err := step.Store(s, out); err != nil {
				return nil, err
			}
		}

		if err := s.scanTrackerRepo.MarkScanCompleted(domain.Id, step.Name()); err != nil {
			return nil, err
		}
	} else {
		logger.Infof("Step %s already completed for domain %s, fetching results from DB", step.Name(), domain.Name)
	}

	results, err := step.Fetch(s)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch results from db %v", err)
	}

	logger.Successf("Task %s completed successfully. Recieved %d results", step.Name(), len(results))

	return results, nil
}
