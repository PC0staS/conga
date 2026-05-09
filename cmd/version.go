package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "1.5.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show CONGA version",
	Long:  "Display the current version of CONGA.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("CONGA v%s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
