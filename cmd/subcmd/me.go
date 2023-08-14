package subcmd

import "github.com/spf13/cobra"

var MeCmd = &cobra.Command{
	Use:     "me",
	Aliases: []string{"user", "u"},
	Short:   "Operating personal data",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var MeInfoCmd = &cobra.Command{
	Use:   "info",
	Short: `display your information`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
