package core

import (
	"sync"

	"github.com/ayuxsec/spike/internal/pkg/scanner/db"
	"github.com/ayuxsec/spike/pkg/config"
)

type ToolStep interface {
	Name() string
	Enabled(*Scanner) bool
	RequiredTools() []string
	Run(*Scanner, []string) ([]string, error)
	Store(*Scanner, []string) error
	Fetch(*Scanner) ([]string, error)
}

type Scanner struct {
	InputDomains    []string
	currentDomain   *db.Domain // internal
	DBPath          string
	doSaveToDB      bool                      // internal
	dbClient        *db.DB                    // internal
	domainRepo      *db.DomainRepository      // internal
	scanTrackerRepo *db.ScanTrackerRepository // internal
	toolsRepo       *db.ToolsRepository       // internal
	ToolsCfg        *config.ToolsConfig
	ScanMode        ScanMode
	cacheResults    map[string][]string // internal
	EventHandler    ScanEventHandler
	wg              sync.WaitGroup // internal
}

type ScannerOutput struct {
	DomainScanned     string
	Subdomains        []string
	CachexResults     []string
	HttpServers       []string
	GauURLs           []string
	KatanaCrawledURLs []string
	NucleiFindings    []string
}

type ScanEventHandler interface {
	OnDomainScanned(out *ScannerOutput, errs []error)
}

// DbScanMode or ManualScanMode
type ScanMode int

const (
	DbScanMode ScanMode = iota
	ManualScanMode
)
