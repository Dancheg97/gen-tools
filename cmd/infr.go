package cmd

import "github.com/spf13/cobra"

var infraCmd = &cobra.Command{
	Use:     "infr",
	Short:   "🐋 Generate all infrastructure in single command",
	Run:     Infr,
	Example: "gen-tools infr",
}

func init() {
	rootCmd.AddCommand(infraCmd)
}

func Infr(cmd *cobra.Command, args []string) {
	Gen(cmd, []string{
		"compose-nginx",
		"compose-gitea",
		"compose-pocketbase",
		"compose-drone",
		"compose-mkdocs",
		"compose-pacman",
	})
}
