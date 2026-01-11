package core

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ayuxsec/spike/internal/pkg/scanner/cli"
	"github.com/ayuxsec/spike/internal/pkg/scanner/db"
	"github.com/ayuxsec/spike/pkg/config"
	"github.com/ayuxsec/spike/pkg/logger"
)

// NewScanner creates and initializes a new Scanner instance.
func NewScanner(
	ToolsCfg *config.ToolsConfig,
	ReporterCfg *config.ReporterConfig,
	inputDomains []string,
	dbPath string,
	scanMode ScanMode,
) (*Scanner, error) {

	if err := cli.WarnIfToolsMissing(); err != nil {
		return &Scanner{}, err
	}

	s := &Scanner{
		InputDomains: inputDomains,
		DBPath:       dbPath,
		ToolsCfg:     ToolsCfg,
		doSaveToDB:   dbPath != "",
		ScanMode:     scanMode,
	}

	if dbPath == "" {
		return s, nil
	}

	// initialize database connection
	dbClient := &db.DB{}
	logger.Debugf("using db path: %s", dbPath)
	if err := dbClient.Connect(dbPath); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	s.dbClient = dbClient

	// initialize repositories
	s.domainRepo = db.NewDomainRepository(dbClient)
	s.scanTrackerRepo = db.NewScanTrackerRepository(dbClient)
	s.toolsRepo = db.NewToolsRepository(dbClient)

	if err := s.domainRepo.CreateTable(); err != nil {
		return nil, fmt.Errorf("failed to create domains table: %v", err)
	}

	if err := s.scanTrackerRepo.CreateTable(); err != nil {
		return nil, fmt.Errorf("failed to create scan_tracker table: %v", err)
	}

	if err := s.toolsRepo.CreateTables(); err != nil {
		return nil, fmt.Errorf("failed to create tools tables: %v", err)
	}

	return s, nil
}

// ScanDomains scans all input domains based on the configured scan mode.
func (s *Scanner) ScanDomains() error {
	var domainsToScan []string // domain names

	switch s.ScanMode {
	case ManualScanMode:
		if len(s.InputDomains) == 0 {
			return errors.New("no input domains provided for manual scan mode")
		}
		domainsToScan = s.InputDomains

		if s.doSaveToDB {
			if err := s.domainRepo.BulkInsert(domainsToScan); err != nil {
				return fmt.Errorf("failed to bulk-insert input domains into database: %v", err)
			}
			logger.Successf("Inserted %d domain(s) into database (duplicates are ignored)", len(s.InputDomains))
		}
	case DbScanMode:
		unscanned, err := s.domainRepo.GetUnscanned()
		if err != nil {
			return fmt.Errorf("failed to fetch unscanned domains from database: %v", err)
		}
		for _, d := range unscanned {
			domainsToScan = append(domainsToScan, d.Name)
		}
	}

	logger.Infof("Starting scan for %d domain(s)...\n", len(domainsToScan))

	var scanErrs []error
	for _, domainName := range domainsToScan {
		_, errs := s.scanDomain(domainName)
		if len(errs) > 0 {
			var b strings.Builder
			b.WriteString("Errors occurred on some domains. Summary below:\n")
			fmt.Fprintf(&b, "Errors while scanning %s:\n", domainName)

			for _, e := range errs {
				if e != nil {
					b.WriteString("• " + e.Error() + "\n") // aggregate all errors into a single error for this domain
				}
			}

			b.WriteString("\n") // separate domain errors by a newline

			scanErrs = append(scanErrs, errors.New(strings.TrimSpace(b.String()))) // append to overall scan errors
		}
	}

	if len(scanErrs) == 0 {
		return nil
	}

	// write all scan errors into a single error
	var out strings.Builder
	for _, e := range scanErrs {
		out.WriteString(e.Error())
		out.WriteString("\n\n") // separate errors by double newlines
	}

	return errors.New(strings.TrimSpace(out.String()))
}

func (s *Scanner) scanDomain(domainName string) (*ScannerOutput, []error) {
	var errs []error
	so := &ScannerOutput{DomainScanned: domainName}

	// ----------------------------------------------------
	// Load Domain
	// ----------------------------------------------------
	d, err := s.domainRepo.SelectByName(domainName)
	if err != nil {
		errs = append(errs, fmt.Errorf("failed to fetch domain: %v", err))
		return so, errs
	}
	s.currentDomain = d

	logger.Infof("=== Starting scan for domain: %s ===", d.Name)

	// ----------------------------------------------------
	// Initialize pipeline result cache
	// ----------------------------------------------------
	s.cacheResults = make(map[string][]string)
	pipeline := s.selectPipeline()
	s.cacheResults["__root_domain__"] = []string{d.Name} // root domain for tools that need it

	// ----------------------------------------------------
	// Execute pipeline sequentially
	// ----------------------------------------------------
	for _, p := range pipeline {
		step := p.Step
		stepName := step.Name()

		fmt.Fprintln(os.Stderr) // trailing newline for better readability

		// Collect correct inputs based on dependencies
		var input []string
		if p.InputFrom != nil {
			for _, dep := range p.InputFrom {
				input = append(input, s.cacheResults[dep]...)
			}
		}

		out, err := s.ExecStep(step, input)
		if err != nil {
			logger.Errorf("%s step failed: %v", stepName, err)
			errs = append(errs, fmt.Errorf("%s step failed: %v", stepName, err))
		}

		// Cache the result for later steps (Nuclei needs multiple)
		s.cacheResults[stepName] = out
	}

	// ----------------------------------------------------
	// Map results into ScannerOutput
	// ----------------------------------------------------
	so.Subdomains = s.cacheResults["subfinder"]
	so.HttpServers = s.cacheResults["httpx"]
	so.GauURLs = s.cacheResults["gau"]
	so.CachexResults = s.cacheResults["cachex"]
	so.KatanaCrawledURLs = s.cacheResults["katana"]
	so.NucleiFindings = append(s.cacheResults["nuclei_generic"], s.cacheResults["nuclei_dast"]...)

	if s.EventHandler != nil {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.EventHandler.OnDomainScanned(so, errs)
		}()
	} else {
		logger.Warn("No ScanEventHandler registered, skipping OnDomainScanned event")
	}

	if isScanned, err := s.checkAllToolsScanned(); err != nil {
		errs = append(errs, fmt.Errorf("failed to verify if all tools scanned: %v", err))
	} else if isScanned {
		if err := s.domainRepo.MarkAsScanned(domainName); err != nil {
			errs = append(errs, fmt.Errorf("failed to mark domain as scanned: %v", err))
		}
	}

	logger.Successf("=== Completed scan for %s ===\n", d.Name)

	return so, errs
}

func (s *Scanner) Close() error {
	s.wg.Wait() // ensure all notifiers are done
	// close db connection
	if s.dbClient != nil {
		return s.dbClient.Close()
	}
	return nil
}
