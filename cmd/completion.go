package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate shell completion script",
	Long: `Generate the autocompletion script for conga for the specified shell.

To load completions:

Bash:
  source <(conga completion bash)

  To load completions for each session, execute once:
  conga completion bash > /etc/bash_completion.d/conga

Zsh:
  source <(conga completion zsh)

  To load completions for each session, execute once:
  conga completion zsh > "${fpath[1]}/_conga"

Fish:
  conga completion fish | source

  To load completions for each session, execute once:
  conga completion fish > ~/.config/fish/completions/conga.fish

PowerShell:
  conga completion powershell | Out-String | Invoke-Expression

  To load completions for every new session, run:
  conga completion powershell > conga.ps1
  and source this file from your PowerShell profile.`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
