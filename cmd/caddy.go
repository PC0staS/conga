package cmd

import (
	"github.com/pc0stas/conga/generators"
	"github.com/spf13/cobra"
)

var caddyCmd = &cobra.Command{
	Use:   "caddy",
	Short: "Caddy configuration generator",
	Long: `Caddy configuration generator.

Interactive flow to create a simple Caddyfile with:
- Main domain
- Optional TLS (provide cert/key paths)
- Multiple paths (reverse_proxy or file_server)

Generates a 'Caddyfile' ready to use.`,
}

var caddyGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Caddy config",
	Long:  "Interactive flow to create a Caddyfile with domain, TLS, and routes.",
	Run: func(cmd *cobra.Command, args []string) {
		generators.GenerateCaddyConfig()
	},
}

func init() {
	rootCmd.AddCommand(caddyCmd)
	caddyCmd.AddCommand(caddyGenerateCmd)
}
