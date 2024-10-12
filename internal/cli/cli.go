package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/loyality7/Isecot/internal/scanner"
)

func Start() {
	fmt.Println("IoT Security Tool - CLI Mode")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("iot> ")
		cmdString, _ := reader.ReadString('\n')
		cmdString = strings.TrimSpace(cmdString)
		args := strings.Fields(cmdString)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "exit":
			fmt.Println("Goodbye!")
			os.Exit(0)
		case "scan":
			scanner.ScanNetwork()
		default:
			fmt.Println("Unknown command. Available commands: scan, exit")
		}
	}
}
