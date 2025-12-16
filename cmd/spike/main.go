package main

import (
	"os"
	"spike/internal/app/spike/cmd"
	"spike/pkg/logger"
)

func main() {
	if err := cmd.App().Run(os.Args); err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
}
