package config

type Config struct {
	Tools    ToolsConfig    `yaml:"tools"`
	Reporter ReporterConfig `yaml:"reporter"`
}

type ToolsConfig struct {
	HTTPX     HTTPXConfig     `yaml:"httpx"`
	Subfinder SubfinderConfig `yaml:"subfinder"`
	Katana    KatanaConfig    `yaml:"katana"`
	Gau       GauConfig       `yaml:"gau"`
	Nuclei    NucleiConfig    `yaml:"nuclei"`
	Cachex    CachexConfig    `yaml:"cachex"`
}

type CachexConfig struct {
	Enabled    bool `yaml:"enabled"`
	Threads    int  `yaml:"threads"`
	CmdTimeout int  `yaml:"cmd_timeout_in_second"`
}

type HTTPXConfig struct {
	Threads     int             `yaml:"threads"`
	TargetPorts HttxPortsConfig `yaml:"ports_to_scan"`
	ScreenShot  bool            `yaml:"screenshot"`
	CmdTimeout  int             `yaml:"cmd_timeout_in_second"`
}

type HttxPortsConfig struct {
	Http  string `yaml:"http"`  // comma separated list of ports
	Https string `yaml:"https"` // comma separated list of ports
}

type SubfinderConfig struct {
	Threads    int `yaml:"threads"`
	ActiveEnum SubsActiveEnumConfig
	Enabled    bool `yaml:"enabled"`
	CmdTimeout int  `yaml:"cmd_timeout_in_second"`
}

type SubsActiveEnumConfig struct {
	Enabled      bool   `yaml:"enabled"`
	WordlistMode string `yaml:"wordlist_mode"` // wordlist mode to use (magic string: large, tiny...)
	DnsxThreads  int    `yaml:"dnsx_threads"`
}

type KatanaConfig struct {
	Enabled            bool   `yaml:"enabled"`
	Threads            int    `yaml:"threads"`
	CrawlDepth         int    `yaml:"crawl_depth"`
	MaxCrawlTime       string `yaml:"max_crawl_time"`
	ParallelismThreads int    `yaml:"parallelism_threads"`
	Headless           bool   `yaml:"headless"`
	NoSandbox          bool   `yaml:"no_sandbox"`
	CmdTimeout         int    `yaml:"cmd_timeout_in_second"`
}

type GauConfig struct {
	Enabled    bool `yaml:"enabled"`
	Threads    int  `yaml:"threads"`
	CmdTimeout int  `yaml:"cmd_timeout_in_second"`
}

type NucleiConfig struct {
	Threads          int                    `yaml:"threads"`
	TemplateSettings NucleiTemplateSettings `yaml:"template_settings"`
	TemplatePaths    NucleiTemplatePaths    `yaml:"template_paths"`
	CmdTimeout       int                    `yaml:"cmd_timeout_in_second"`
}

type NucleiTemplatePaths struct {
	Generic NucleiTemplatePathSet `yaml:"generic"`
	Dast    NucleiTemplatePathSet `yaml:"dast"`
}

type NucleiTemplatePathSet struct {
	Include         []string `yaml:"include"`
	Exclude         []string `yaml:"exclude"`
	ExcludeSeverity []string `yaml:"exclude_severity"`
}

type NucleiTemplateSettings struct {
	Generic  bool `yaml:"generic"`
	Dast     bool `yaml:"dast"`
	Headless bool `yaml:"headless"`
}

type ReporterConfig struct {
	Telegram TelegramConfig `yaml:"telegram"`
}

type TelegramConfig struct {
	Enabled  bool   `yaml:"enabled"`
	BotToken string `yaml:"bot_token"`
	ChatID   int    `yaml:"chat_id"`
	Timeout  int    `yaml:"request_timeout"`
}
