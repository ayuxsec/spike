package cmd

import (
	"github.com/ayuxsec/spike/internal/pkg/scanner/core"

	"github.com/spf13/cobra"
)

func NewScanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Run a manual scan (append prefix *. on your root domain for wildcards)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runScan(core.ManualScanMode)
		},
	}

	cmd.Flags().StringVarP(
		&domain, "domain", "d", "", "Target domain to scan",
	)
	cmd.Flags().StringVarP(
		&domainList, "domain-list", "l", "", "File containing domains to scan",
	)

	cmd.AddCommand(NewScanDBCmd())

	return cmd
}
