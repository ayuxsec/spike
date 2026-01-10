package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ayuxsec/spike/internal/pkg/scanner/core"
	"github.com/ayuxsec/spike/pkg/config"
	"github.com/ayuxsec/spike/pkg/spike"
)

func runScan(scanMode core.ScanMode) error {
	var cfg config.Config

	if loadCfgPath != "" {
		var err error
		cfg, err = config.LoadCfg(loadCfgPath)
		if err != nil {
			return fmt.Errorf("failed to load config from the specified path: %v", err)
		}
	} else {
		cfg = *config.DefaultConfig()
	}

	var domains []string
	if scanMode == core.ManualScanMode {
		var err error
		domains, err = collectDomains()
		if err != nil {
			return err
		}
	}

	scanner, err := spike.NewScanner(
		spike.OptionWithConfig(&cfg),
		spike.OptionWithDBPath(dbPath),
		spike.OptionWithInputDomains(domains),
	)

	if err != nil {
		return fmt.Errorf("failed to create scanner: %v", err)
	}

	return scanner.Run()
}

func collectDomains() ([]string, error) {
	var domains []string

	if domain != "" {
		return []string{domain}, nil
	}

	if domainList != "" {
		list, err := fileToSlice(domainList)
		if err != nil {
			return nil, fmt.Errorf("failed to read domains from file: %w", err)
		}
		domains = append(domains, list...)
	}

	stat, err := os.Stdin.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get stdin info: %w", err)
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				domains = append(domains, line)
			}
		}
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error reading piped input: %w", err)
		}
	}

	if len(domains) == 0 {
		return nil, fmt.Errorf("no domains provided. Use -d, -l, or pipe input")
	}

	return domains, nil
}
