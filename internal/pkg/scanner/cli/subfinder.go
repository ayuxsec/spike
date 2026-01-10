// This file contains tools related to subdomain discovery on domains.
package cli

import (
	"strconv"

	"github.com/ayuxsec/spike/pkg/config"
)

// RunSubfinder executes the subfinder tool for the given domain with specified threads
func RunSubfinder(domain string, s *config.SubfinderConfig) ([]string, error) {
	args := []string{"-d", domain, "-all", "-t", strconv.Itoa(s.Threads)}
	return RunCommand("subfinder", args, s.CmdTimeout)
}
