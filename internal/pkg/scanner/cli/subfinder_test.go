package cli

import (
	"testing"

	"github.com/ayuxsec/spike/internal/pkg/scanner/db"
	"github.com/ayuxsec/spike/pkg/config"
)

func TestRunSubfinder(t *testing.T) {
	s := config.DefaultConfig().Tools.Subfinder

	subdomains, err := RunSubfinder("example.com", &s)
	if err != nil {
		t.Fatalf("RunSubfinder failed: %v", err)
	}

	// NOTE: Real subfinder must be installed for non-zero results.
	// If not installed, this test will naturally fail.
	if len(subdomains) == 0 {
		t.Logf("No subdomains returned (expected if subfinder binary not available)")
	}
}

func TestSubdomainsBulkInsert(t *testing.T) {
	subs := []string{"sub1.example.com", "sub2.example.com"}

	// Setup database
	dbClient := &db.DB{}
	if err := dbClient.Connect("test.db"); err != nil {
		t.Fatalf("Failed to connect to DB: %v", err)
	}

	// Initialize repos
	domainRepo := db.NewDomainRepository(dbClient)
	toolsRepo := db.NewToolsRepository(dbClient)

	// Create tables
	if err := domainRepo.CreateTable(); err != nil {
		t.Fatalf("Failed to create domains table: %v", err)
	}

	if err := toolsRepo.CreateTables(); err != nil {
		t.Fatalf("Failed to create tools tables: %v", err)
	}

	// Insert test domain
	if err := domainRepo.Insert("example.com"); err != nil {
		t.Fatalf("Failed to insert domain: %v", err)
	}

	// Fetch domain ID
	domains, err := domainRepo.GetAll()
	if err != nil || len(domains) == 0 {
		t.Fatalf("Failed to retrieve domain: %v", err)
	}

	domain, err := domainRepo.SelectByName("example.com")
	if err != nil {
		t.Fatalf("SelectByName failed: %v", err)
	}

	// Insert subdomains using toolsRepo
	if err := toolsRepo.Subfinder.BulkInsert(domain.Id, subs); err != nil {
		t.Fatalf("BulkInsert failed: %v", err)
	}

	t.Logf("Successfully inserted subdomains: %v", subs)
}
