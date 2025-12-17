package cmd

import (
	"bufio"
	"fmt"
	"os"
	"spike/internal/scanner/core"
	"spike/pkg/config"
	"spike/pkg/logger"
	"spike/pkg/spike"
	"spike/pkg/version"
	"strings"

	"github.com/urfave/cli/v2"
)

func App() *cli.App {
	return &cli.App{
		Flags:                 buildFlags(),
		CustomAppHelpTemplate: buildHelpMessage(),
		Commands:              buildCommands(),
		Version:               version.Version,
	}
}

func Run(scanMode core.ScanMode) error {
	PrintBanner()
	logger.DisableDebug = !verbose

	logger.Debugf("using config file: %s", cfgPath)

	if err := ensureAppDirExists(); err != nil {
		return fmt.Errorf("failed to create app directory: %v", err)
	}

	if err := config.LoadOrCreateConfig(cfgPath, false); err != nil {
		return fmt.Errorf("failed to load config file '%s': %v", cfgPath, err)
	}
	var domains []string

	if scanMode == core.ManualScanMode {
		var err error
		domains, err = collectDomains()
		if err != nil {
			return err
		}
	}
	// create spike wrapper
	scanner, err := spike.New(
		&config.Cfg.ToolsConfig,
		&config.Cfg.Reporter,
		domains,
		dbPath,
		scanMode,
	)
	if err != nil {
		return fmt.Errorf("failed to create scanner: %v", err)
	}

	// run engine
	return scanner.Run()
}

func collectDomains() ([]string, error) {
	var domains []string

	// If single domain flag used
	if domain != "" {
		return []string{domain}, nil
	}

	// If domain file provided
	if domainsFilePath != "" {
		list, err := fileToSlice(domainsFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read domains from file: %w", err)
		}
		domains = append(domains, list...)
	}

	// Check for piped input
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
		return nil, fmt.Errorf("no domains provided. Use -d or -dl flag, or provide piped input")
	}

	return domains, nil
}
