package cmd

import (
	"github.com/pc0stas/conga/generators"
	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker configuration generator",
	Long: `Docker Compose configuration generator.

Interactive flow to create docker-compose.yml with:
- Custom services
- Image selection
- Port mapping
- Volume mounting
- Environment variables
- Environment files

Generates a 'docker-compose.yml' file ready to use.`,
}

var dockerGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Docker Compose config",
	Long:  "Interactive flow to create a docker-compose.yml with services, ports, volumes, and environment variables.",
	Run: func(cmd *cobra.Command, args []string) {
		generators.GenerateDockerCompose()
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
	dockerCmd.AddCommand(dockerGenerateCmd)
}
