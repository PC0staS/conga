package main

import "fmt"

func PrintMainHelp() {
	fmt.Println(`
🔧 CONGA - Config Generator App
================================

Usage: conga <service> <command>

Available services:
  nginx     - Nginx configuration generator

Examples:
  conga nginx generate     - Generate Nginx config
  conga nginx help         - Show Nginx help
  conga help               - Show this help`)
}