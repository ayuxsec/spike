package cmd

import (
	"spike/internal/scanner/core"

	"github.com/urfave/cli/v2"
)

func buildCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "scan",
			Usage: "Run a scan",
			Flags: buildFlags(),

			// spike scan -d example.com
			Action: func(c *cli.Context) error {
				return Run(core.ManualScanMode)
			},

			// spike scan db --db ./spike.db
			Subcommands: []*cli.Command{
				{
					Name:  "db",
					Usage: "Run DB scan mode",
					Flags: buildFlags(),

					// spike scan db
					Action: func(c *cli.Context) error {
						return Run(core.DbScanMode)
					},
				},
			},
		},
	}
}
