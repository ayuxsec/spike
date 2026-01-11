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
					Http:  "80,8080,8000,8008,8888,3000,5000,9000,81,82,83,84,591,2082,2086,2095,10000",
					Https: "443,8443,9443,5001,3001,8001,8081,2083,2087,2096,10001,10443,10444",
				},
				ScreenShot: false,
				CmdTimeout: GlobalCmdTimeout,
			},
			Subfinder: SubfinderConfig{
				Threads: 10,
				ActiveEnum: SubsActiveEnumConfig{
					Enabled:      true,
					WordlistMode: "tiny",
					DnsxThreads:  200,
				},
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
