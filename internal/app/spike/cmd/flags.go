package cmd

import (
	"path/filepath"
	"spike/pkg/config"

	"github.com/urfave/cli/v2"
)

var (
	domain          string
	domainsFilePath string
	dbPath          string = filepath.Join(config.DefaultAppDir, "spike.db")    // default database path
	cfgPath         string = filepath.Join(config.DefaultAppDir, "config.yaml") // default config path
	verbose         bool   = true
)

func buildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "domain",
			Aliases:     []string{"d"},
			Destination: &domain,
		},
		&cli.StringFlag{
			Name:        "domain-list",
			Aliases:     []string{"dl"},
			Destination: &domainsFilePath,
		},
		&cli.StringFlag{
			Name:        "database",
			Aliases:     []string{"db"},
			Destination: &dbPath,
			Value:       dbPath,
		},
		&cli.BoolFlag{
			Name:        "debug",
			Aliases:     []string{"verbose"},
			Destination: &verbose,
			Value:       verbose,
		},
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Destination: &cfgPath,
			Value:       cfgPath,
		},
	}
}
