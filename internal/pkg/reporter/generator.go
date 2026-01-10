package reporter

import (
	"fmt"
	"strings"

	scanner "github.com/ayuxsec/spike/internal/pkg/scanner/core"
	"github.com/ayuxsec/spike/internal/pkg/scanner/db"
)

// GenerateDBReport generates a report from the database.
func GenerateDBReport(domainRepo *db.DomainRepository) (string, error) {
	var b strings.Builder
	allDomains, err := domainRepo.GetAll()
	if err != nil {
		return "", fmt.Errorf("failed to select all domains from databbase: %v", err)
	}
	scannedDomains, err := domainRepo.GetScanned()
	if err != nil {
		return "", fmt.Errorf("failed to select scanned domains from database: %v", err)
	}
	fmt.Fprintf(&b, "Scanned domains: %d/%d\n", len(scannedDomains), len(allDomains))

	return b.String(), nil
}

// GenerateScanReport generates a report from the scanner output.
func GenerateScanReport(so *scanner.ScannerOutput) (string, error) {
	var b strings.Builder

	fmt.Fprintf(&b, "Scan Report for %s\n", so.DomainScanned)
	fmt.Fprintf(&b, "Subdomains found: %d\n", len(so.Subdomains))
	fmt.Fprintf(&b, "Live hosts found: %d\n", len(so.HttpServers))
	fmt.Fprintf(&b, "Crawled URLs found: %d\n", len(append(so.KatanaCrawledURLs, so.GauURLs...)))
	fmt.Fprintf(&b, "Nuclei Scan Vulnerabilities found: %d\n", len(so.NucleiFindings))
	fmt.Fprintf(&b, "Cache Poisoning Vulnerabilitie found: %d\n", len(so.CachexResults))

	// if len(so.NucleiFindings) > 0 {
	// 	b.WriteString("\nNuclei Scan Vulnerabilities:\n")
	// 	for _, nucleiOutput := range so.NucleiFindings {
	// 		b.WriteString(nucleiOutput)
	// 		b.WriteByte('\n')
	// 	}
	// }

	// if len(so.CachexResults) > 0 {
	// 	b.WriteString("\nCachex Results:\n")
	// 	for _, cachexOutput := range so.CachexResults {
	// 		b.WriteString(cachexOutput)
	// 		b.WriteByte('\n')
	// 	}
	// }

	return b.String(), nil
}

func generateErrsReport(domainName string, errs []error) string {
	if len(errs) == 0 {
		return ""
	}
	var b strings.Builder
	fmt.Fprintf(&b, "Errors occurred during scan of %s:\n", domainName)
	for _, e := range errs {
		if e != nil {
			b.WriteString("• " + e.Error() + "\n")
		}
	}
	return b.String()
}
