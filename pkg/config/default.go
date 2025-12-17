package config

import (
	"os"
	"path/filepath"
)

var homeDir = os.Getenv("HOME")
var appName = "spike"
var DefaultAppDir = filepath.Join(homeDir, appName)

func DefaultConfig() *Config {
	//logger.Debugf("Using $HOME directory as: %s", homeDir)
	return &Config{
		ToolsConfig: ToolsConfig{
			HTTPX: HTTPXConfig{
				Threads:    25,
				ScreenShot: false,
			},
			Subfinder: SubfinderConfig{
				Threads: 10,
				Enabled: true,
			},
			Gau: GauConfig{
				Enabled: true,
				Threads: 10,
			},
			Katana: KatanaConfig{
				Enabled:            true,
				Threads:            10,
				CrawlDepth:         3,
				MaxCrawlTime:       "10m",
				ParallelismThreads: 10,
				Headless:           false,
				NoSandbox:          false,
			},
			Cachex: CachexConfig{ // check ~/cachex/config.yaml for better customization
				Enabled: true,
				Threads: 10,
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
