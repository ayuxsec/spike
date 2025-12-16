package cmd

import (
	"fmt"
	"os"
	"spike/pkg/version"
)

var banner = `
               _ __      
   _________  (_) /_____ 
  / ___/ __ \/ / //_/ _ \
 (__  ) /_/ / / ,< /  __/
/____/ .___/_/_/|_|\___/ 
    /_/                  
`

func PrintBanner() {
	fmt.Fprint(os.Stderr, banner)
	fmt.Fprint(os.Stderr, "    "+version.Version+" with <3 by @ayuxsec\n\n")
}
