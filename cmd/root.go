package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "conga",
	Short: "CONGA - Config Generator App",
	Long: `🔧 CONGA - Config Generator App
================================

Interactive CLI tool for generating configuration files for
Nginx, Docker, Caddy, Apache, and WireGuard.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
