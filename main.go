package main

import (
	"fmt"
	"os"

	"github.com/pc0stas/conga/generators"
)
var Version = "0.1.0"

func main() {
	// Check if there are arguments
	if len(os.Args) < 2 {
		PrintMainHelp()
		os.Exit(0)
	}

	service := os.Args[1]
	command := "help"
	if len(os.Args) > 2 {
		command = os.Args[2]
	}

	// Dispatcher based on service
	switch service {
	case "nginx":
		generators.HandleNginx(command)
	case "version":
		fmt.Printf("CONGA v%s\n", Version)

	case "help":
		PrintMainHelp()
	default:
		fmt.Printf("Unknown service: %s\n", service)
		os.Exit(1)
	}
}