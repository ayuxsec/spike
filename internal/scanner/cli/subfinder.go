// This file contains tools related to subdomain discovery on domains.
package cli

import (
	"strconv"
)

// RunSubfinder executes the subfinder tool for the given domain with specified threads
func RunSubfinder(domain string, Threads int) ([]string, error) {
	args := []string{"-d", domain, "-all", "-t", strconv.Itoa(Threads)}
	return RunCommand("subfinder", args)
}
