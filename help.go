package main

import "fmt"

func PrintMainHelp() {
	fmt.Println(`
🔧 CONGA - Config Generator App
================================

Usage: conga <service> <command>

Available services:
  nginx     - Nginx configuration generator
  docker    - Docker configuration generator
  version   - Show CONGA version
  help      - Show this help message

Examples:
  conga nginx generate     - Generate Nginx config
  conga nginx help         - Show Nginx help
  conga docker generate    - Generate Docker config
  conga docker help        - Show Docker help
  conga version            - Show CONGA version
  conga help               - Show this help`)
}