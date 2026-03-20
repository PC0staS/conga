package main

import "fmt"

func PrintMainHelp() {
	fmt.Println(`
🔧 CONGA - Config Generator App
================================

Usage: conga <service> <command>

For a full list of available services and commands, visit:
  https://github.com/PC0staS/conga

Examples:
  conga nginx generate     - Generate Nginx config
  conga nginx help         - Show Nginx help
  conga help               - Show this help`)
}