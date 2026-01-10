package cmd

import (
	"fmt"

	"github.com/ayuxsec/spike/pkg/spike"

	"github.com/spf13/cobra"
)

func NewShellCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shell",
		Short: "Drop into spike REPL shell",
		RunE: func(cmd *cobra.Command, args []string) error {
			if dbPath == "" {
				return fmt.Errorf("please specify a SQLite database path using --db")
			}
			return spike.NewREPLShell(dbPath)
		},
	}
	return cmd
}
