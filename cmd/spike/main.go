package main

import (
	"os"

	"github.com/ayuxsec/spike/internal/app/spike/cmd"
	"github.com/ayuxsec/spike/pkg/logger"
)

func main() {
	if err := ExecuteCmd(); err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
}

func ExecuteCmd() error {
	cmd.PrintBanner()
	cmd := cmd.NewRootCmd()
	cmd.SilenceErrors = true
	return cmd.Execute()
}
