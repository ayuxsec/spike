// This file contains tools related to subdomain discovery on domains.
package cli

import (
	"fmt"
	"strconv"

	"github.com/ayuxsec/spike/pkg/config"
)

// RunSubfinder executes the subfinder tool for the given domain with specified threads
func RunSubfinder(domain string, s *config.SubfinderConfig) ([]string, error) {
	args := []string{"-d", domain, "-all", "-t", strconv.Itoa(s.Threads)}
	subs, err := RunCommand("subfinder", args, s.CmdTimeout)
	if err != nil {
		return nil, fmt.Errorf("subfinder failed: %v", err)
	}

	// todo: run resub in seperate goroutine.
	var fuzzedSubs []string
	if s.ActiveEnum.Enabled {
		fuzzedSubs, err = runResub(domain, s.CmdTimeout,
			s.ActiveEnum.WordlistMode, s.ActiveEnum.DnsxThreads)
		if err != nil {
			return nil, fmt.Errorf("resub failed: %v", err)
		}
	}

	return append(subs, fuzzedSubs...), nil
}

// runResub creates fuzzable subdomains permutations via replacing FUZZ with nokovo subdomains wordlist
func runResub(domain string, cmdTimeout int, resubWordlistMode string, dnsxThreads int) ([]string, error) {
	fuzzableSubs, err := RunCommand("resub", []string{"FUZZ." + domain, "-m", resubWordlistMode}, cmdTimeout)
	if err != nil {
		return nil, fmt.Errorf("failed to create fuzzable urls: %v", err)
	}
	aliveSubs, err := RunCommandWithStdinInput("dnsx", []string{"-t", strconv.Itoa(dnsxThreads)}, fuzzableSubs, cmdTimeout)
	if err != nil {
		return nil, fmt.Errorf("dnsx failed: %v", err)
	}
	return aliveSubs, nil
}
