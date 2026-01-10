package cmd

import (
	"fmt"
	"os"

	"github.com/ayuxsec/spike/pkg/version"
)

var Banner = `
               _ __      
   _________  (_) /_____ 
  / ___/ __ \/ / //_/ _ \
 (__  ) /_/ / / ,< /  __/
/____/ .___/_/_/|_|\___/ 
    /_/                  
`

func PrintBanner() {
	fmt.Fprint(os.Stderr, Banner)
	fmt.Fprint(os.Stderr, "    "+version.String()+" with <3 by @ayuxsec\n\n")
}
