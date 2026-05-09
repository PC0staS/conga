package cmd

import (
	"github.com/pc0stas/conga/generators"
	"github.com/spf13/cobra"
)

var nginxCmd = &cobra.Command{
	Use:   "nginx",
	Short: "Nginx configuration generator",
	Long: `Nginx configuration generator.

Interactive flow to create Nginx configuration with:
- Main domain
- HTTPS/SSL
- Multiple paths (proxy or static files)
- Standard headers
- Syntax validation

Generates a 'default.conf' file ready to use.`,
}

var nginxGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Nginx config",
	Long:  "Interactive flow to create Nginx configuration with domain, HTTPS, routes, and more.",
	Run: func(cmd *cobra.Command, args []string) {
		generators.GenerateNginxConfig()
	},
}

func init() {
	rootCmd.AddCommand(nginxCmd)
	nginxCmd.AddCommand(nginxGenerateCmd)
}
