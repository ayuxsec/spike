package shell

import (
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/ayuxsec/spike/internal/pkg/scanner/db"
)

// handleCommand handles the command given by the user in the REPL shell
func handleCommand(ctx *Context, args []string) error {
	var output strings.Builder
	const padding = 3

	domainRepo := db.NewDomainRepository(ctx.DB)
	toolsRepo := db.NewToolsRepository(ctx.DB)
	switch args[0] {
	case "help":
		fmt.Fprint(&output, helpString())

	case "domains":
		domains, err := domainRepo.GetAll()
		if err != nil {
			return fmt.Errorf("failed to get domains from database: %v", err)
		}
		w := tabwriter.NewWriter(&output, 0, 0, padding, ' ', 0)
		defer w.Flush()
		fmt.Fprintln(w, "domain\tis_scanned\tcreated_at\t")
		for _, d := range domains {
			fmt.Fprintf(w, "%s\t%s\t%s\t\n", d.Name, strconv.FormatBool(d.IsScanned), d.CreatedAt.Format("2006-01-02 15:04:05"))
		}

	case "select":
		if len(args) != 2 {
			return fmt.Errorf("usage: select <domain>")
		}
		d, err := domainRepo.SelectByName(args[1])
		if err != nil {
			return err
		}
		ctx.Domain = d
		fmt.Println("selected:", d.Name)

	// todo: handle errors
	case "subs":
		data, _ := toolsRepo.Subfinder.Fetch(ctx.Domain.Id)
		fmt.Fprint(&output, strings.Join(data, "\n"))

	case "httpx":
		data, _ := toolsRepo.Httpx.Fetch(ctx.Domain.Id)
		fmt.Fprint(&output, strings.Join(data, "\n"))

	case "uro":
		data, _ := toolsRepo.Uro.Fetch(ctx.Domain.Id)
		fmt.Fprint(&output, strings.Join(data, "\n"))

	case "nuclei":
		data, _ := toolsRepo.Nuclei.Fetch(ctx.Domain.Id)
		fmt.Fprint(&output, strings.Join(data, "\n"))

	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}

	// if user wants to pipe commands allow so
	// todo: output from commands is given after exec completes writing to buffer instead of continous writing
	if len(args) > 1 {
		if strings.HasPrefix(args[1], "|") {
			cmd := exec.Command(args[2], args[3:]...)
			stdinPipe, err := cmd.StdinPipe()
			if err != nil {
				return fmt.Errorf("failed to create stdin pipe: %v", err)
			}
			go func() {
				defer stdinPipe.Close()
				io.WriteString(stdinPipe, output.String())
			}()
			output, err := cmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("failed to run command %s: %v", args[2], err)
			}
			fmt.Printf("%s", output)
		}
	} else {
		fmt.Println(output.String())
	}
	return nil
}

func helpString() string {
	var b strings.Builder
	fmt.Fprint(&b, "\nAvailable Commands:\n")
	for k, v := range Cmds {
		fmt.Fprintf(&b, "  %-10s %s\n", k, v)
	}

	return strings.TrimSuffix(b.String(), "\n")
}
