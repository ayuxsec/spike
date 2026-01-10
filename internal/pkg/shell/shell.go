package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ayuxsec/spike/internal/pkg/scanner/db"
)

func NewREPLShell(dbPath string) error {
	printBanner()

	database := &db.DB{}
	if err := database.Connect(dbPath); err != nil {
		return err
	}
	defer database.Close()

	ctx := &Context{DB: database}

	reader := bufio.NewScanner(os.Stdin)

	for {
		var domainName *string
		if ctx.Domain != nil {
			domainName = &ctx.Domain.Name
		}

		printPrompt(domainName)

		if !reader.Scan() {
			break
		}

		args := parse(reader.Text())
		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		if isExit(cmd) {
			printExitMsg()
			break
		}

		if err := handleCommand(ctx, args); err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func isExit(cmd string) bool {
	return cmd == "exit" || cmd == "quit" || cmd == "q"
}

func parse(input string) []string {
	return strings.Fields(strings.TrimSpace(input))
}
