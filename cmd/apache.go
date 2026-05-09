package cmd

import (
	"github.com/pc0stas/conga/generators"
	"github.com/spf13/cobra"
)

var apacheCmd = &cobra.Command{
	Use:   "apache",
	Short: "Apache configuration generator",
	Long: `Apache VirtualHost configuration generator.

Interactive flow to create an Apache VirtualHost configuration with:
- Main domain
- Optional TLS/SSL
- Multiple paths (ProxyPass or Alias)

Generates an 'apache.conf' file ready to use.`,
}

var apacheGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Apache config",
	Long:  "Interactive flow to create an Apache VirtualHost with domain, TLS, and routes.",
	Run: func(cmd *cobra.Command, args []string) {
		generators.GenerateApacheConfig()
	},
}

func init() {
	rootCmd.AddCommand(apacheCmd)
	apacheCmd.AddCommand(apacheGenerateCmd)
}
