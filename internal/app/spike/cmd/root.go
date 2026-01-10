// internal/app/spike/cmd/root.go
package cmd

import (
	"github.com/ayuxsec/spike/pkg/config"
	"github.com/ayuxsec/spike/pkg/logger"
	"github.com/ayuxsec/spike/pkg/version"

	"github.com/spf13/cobra"
)

var (
	domain        string
	domainList    string
	dbPath        string = "spike.db"
	loadCfgPath   string
	createCfgPath string
	verbose       bool = true
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "spike",
		Short: "Advanced Stateful web scanner",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.DisableDebug = !verbose
			if createCfgPath != "" {
				return config.CreateDefaultCfg(createCfgPath)
			}
			return cmd.Help()
		},
	}

	rootCmd.PersistentFlags().StringVar(
		&dbPath, "db", "spike.db", "Path to SQLite database",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&verbose, "debug", "", verbose, "Enable debug logs",
	)
	rootCmd.PersistentFlags().StringVar(
		&loadCfgPath, "load-config", "", "Path to YAML config file for loading spike configurations",
	)

	rootCmd.PersistentFlags().StringVar(
		&createCfgPath, "write-config", createCfgPath, "Path to write the default YAML config file",
	)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print version and exit",
		Run: func(_ *cobra.Command, args []string) {
			logger.Infof("current spike version in use: %s", version.String())
		},
	})

	rootCmd.AddCommand(NewScanCmd())
	rootCmd.AddCommand(NewShellCmd())

	return rootCmd
}
