package config

import (
	"os"
	"path/filepath"
)

var homeDir = os.Getenv("HOME")
var appName = "spike"
var DefaultAppDir = filepath.Join(homeDir, appName)

const GlobalCmdTimeout = 900

func DefaultConfig() *Config {
	//logger.Debugf("Using $HOME directory as: %s", homeDir)
	return &Config{
		Tools: ToolsConfig{
			HTTPX: HTTPXConfig{
				Threads: 25,
				TargetPorts: HttxPortsConfig{
					Http: []string{
						"80",                   // Standard HTTP
						"8080",                 // Common alternate HTTP
						"8000",                 // Dev servers
						"8008",                 // Proxies / alt HTTP
						"8888",                 // Dashboards / dev tools
						"3000",                 // Node / frontend dev
						"5000",                 // Flask / APIs
						"9000",                 // Internal apps / PHP frontends
						"81", "82", "83", "84", // Extra HTTP ports
						"591",   // FileMaker web
						"2082",  // cPanel HTTP
						"2086",  // WHM HTTP
						"2095",  // Webmail HTTP
						"10000", // Webmin (often HTTP)
					},
					Https: []string{
						"443",  // Standard HTTPS
						"8443", // Admin panels, dashboards
						"9443", // Enterprise / management consoles
						// Dev / API servers over TLS (Flask, FastAPI, Node, Go, etc.)
						"5001",  // HTTPS version of Flask (5000)
						"3001",  // HTTPS Node / frontend
						"8001",  // HTTPS dev servers
						"8081",  // HTTPS alternate
						"2083",  // cPanel HTTPS
						"2087",  // WHM HTTPS
						"2096",  // Webmail HTTPS
						"10001", // Webmin over HTTPS (some setups)
						"10443", // Alternate TLS
						"10444", // Alternate TLS
					},
				},
				ScreenShot: false,
				CmdTimeout: GlobalCmdTimeout,
			},
			Subfinder: SubfinderConfig{
				Threads:    10,
				Enabled:    true,
				CmdTimeout: GlobalCmdTimeout,
			},
			Gau: GauConfig{
				Enabled:    true,
				Threads:    10,
				CmdTimeout: GlobalCmdTimeout,
			},
			Katana: KatanaConfig{
				Enabled:            true,
				Threads:            10,
				CrawlDepth:         3,
				MaxCrawlTime:       "10m",
				ParallelismThreads: 10,
				Headless:           false,
				NoSandbox:          false,
				CmdTimeout:         GlobalCmdTimeout,
			},
			Cachex: CachexConfig{ // check ~/cachex/config.yaml for better customization
				Enabled:    true,
				Threads:    10,
				CmdTimeout: GlobalCmdTimeout,
			},
			Nuclei: NucleiConfig{
				Threads: 25,
				TemplatePaths: NucleiTemplatePaths{
					Generic: NucleiTemplatePathSet{
						Include: []string{
							"http/",
							"cloud/",
							"javascript/",
							"dns/",
							"ssl/",
							"network/",
							"http/cves/2024",
							"http/cves/2023",
							"http/cves/2022",
							"http/cves/2021",
							"http/cves/2020",
						},
						Exclude: []string{
							"http/cves/", // exclude all other CVE templates
						},
						ExcludeSeverity: []string{"info"},
					},
					Dast: NucleiTemplatePathSet{
						Include:         []string{"dast/"},
						Exclude:         []string{""},
						ExcludeSeverity: []string{"info"},
					},
				},
				TemplateSettings: NucleiTemplateSettings{
					Generic:  true,
					Dast:     true,
					Headless: false,
				},
				CmdTimeout: GlobalCmdTimeout,
			},
		},
		Reporter: ReporterConfig{
			Telegram: TelegramConfig{
				Enabled:  false,
				BotToken: "",
				ChatID:   000000000,
				Timeout:  10,
			},
		},
	}
}
