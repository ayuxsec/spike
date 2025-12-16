package config

type Config struct {
	ToolsConfig ToolsConfig    `yaml:"tools"`
	Reporter    ReporterConfig `yaml:"reporter"`
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
	Enabled bool `yaml:"enabled"`
	Threads int  `yaml:"threads"`
}

type HTTPXConfig struct {
	Threads    int  `yaml:"threads"`
	ScreenShot bool `yaml:"screenshot"`
}

type SubfinderConfig struct {
	Threads int  `yaml:"threads"`
	Enabled bool `yaml:"enabled"`
}

type KatanaConfig struct {
	Enabled            bool   `yaml:"enabled"`
	Threads            int    `yaml:"threads"`
	CrawlDepth         int    `yaml:"crawl_depth"`
	MaxCrawlTime       string `yaml:"max_crawl_time"`
	ParallelismThreads int    `yaml:"parallelism_threads"`
	Headless           bool   `yaml:"headless"`
	NoSandbox          bool   `yaml:"no_sandbox"`
}

type GauConfig struct {
	Enabled bool `yaml:"enabled"`
	Threads int  `yaml:"threads"`
}

type NucleiConfig struct {
	Threads          int                    `yaml:"threads"`
	TemplatePaths    NucleiTemplatePaths    `yaml:"template_paths"`
	TemplateSettings NucleiTemplateSettings `yaml:"template_settings"`
}

type NucleiTemplatePaths struct {
	Generic NucleiTemplatePathSet `yaml:"generic"`
	Dast    NucleiTemplatePathSet `yaml:"dast"`
}

type NucleiTemplatePathSet struct {
	Include []string `yaml:"include"`
	Exclude []string `yaml:"exclude"`
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
	Timeout  int    `yaml:"timeout"`
}
