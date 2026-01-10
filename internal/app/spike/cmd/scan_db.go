package cmd

import (
	"github.com/ayuxsec/spike/internal/pkg/scanner/core"

	"github.com/spf13/cobra"
)

func NewScanDBCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "db",
		Short: "Run scan using domains from the database and store results back into it",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runScan(core.DbScanMode)
		},
	}
}
