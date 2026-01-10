package shell

import (
	"fmt"

	"github.com/ayuxsec/spike/pkg/logger"
)

var Cmds = map[string]string{
	"help":    "Show this help",
	"domains": "List all domains",
	"select":  "Select a domain: select <domain>",
	"subs":    "List subdomains",
	"httpx":   "List live hosts",
	"uro":     "List crawled URLs (active + passive)",
	"nuclei":  "List nuclei findings",
	"exit":    "Exit shell",
	"quit":    "Exit shell",
}

// todo: implement this
func printBanner() {
	fmt.Print(banner)
}

func printPrompt(domain *string) {
	if domain != nil {
		fmt.Printf("\nspike(%s)> ", *domain)
	} else {
		fmt.Print("\nspike> ")
	}
}

func printExitMsg() {
	logger.MagnetaColor.Println("see you space cowboy!")
}
