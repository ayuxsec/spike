package cmd

func buildHelpMessage() string {
	return `spike - simple and fast web scanner

USAGE:
spike scan [flags]        run manual scan
spike scan db [flags]     scan unscanned domains from database

FLAGS:
-h, --help            Show help information
-v, --version         Show application version
-d, --domain          Target domain to scan
-dl, --domain-list    File containing domains to scan (one per line)
-db, --database       SQLite database path
-c,  --config         Path to YAML config file
-debug, --verbose     Enable debug logs
`
}
