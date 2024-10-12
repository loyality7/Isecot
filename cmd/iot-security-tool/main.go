package main

import (
	"os"

	"github.com/loyality7/Isecot/internal/cli"
	"github.com/loyality7/Isecot/internal/web"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		web.StartServer()
	} else {
		cli.Start()
	}
}
