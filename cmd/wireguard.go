package cmd

import (
	"fmt"

	"github.com/pc0stas/conga/generators"
	"github.com/spf13/cobra"
)

var wireguardCmd = &cobra.Command{
	Use:   "wireguard",
	Short: "WireGuard configuration generator",
	Long: `WireGuard configuration generator.

Interactive flow to create a WireGuard configuration with:
- Interface name
- Private key
- Address
- DNS
- Listen port
- Peers

Generates a '<interface>.conf' file ready to use.`,
}

var wireguardGenerateCmd = &cobra.Command{
	Use:       "generate [server|client]",
	Short:     "Generate WireGuard config (server or client)",
	Long:      "Interactive flow to create a WireGuard configuration for server or client mode.",
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgs: []string{"server", "client"},
	Run: func(cmd *cobra.Command, args []string) {
		mode := args[0]
		if mode != "server" && mode != "client" {
			fmt.Printf("Unknown mode: %s\n", mode)
			return
		}
		generators.GenerateWireGuardConfig(mode)
	},
}

func init() {
	rootCmd.AddCommand(wireguardCmd)
	wireguardCmd.AddCommand(wireguardGenerateCmd)
}
