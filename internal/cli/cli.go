package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/loyality7/Isecot/internal/scanner"
	"github.com/loyality7/Isecot/internal/device"
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
			devices := scanner.ScanNetwork()
			for _, d := range devices {
				info, err := device.GetDeviceInfo(d.IP)
				if err == nil {
					fmt.Println(info)
				} else {
					fmt.Println(d)
				}
			}
		default:
			fmt.Println("Unknown command. Available commands: scan, exit")
		}
	}
}
