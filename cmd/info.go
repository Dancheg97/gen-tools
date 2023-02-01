package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "📝 Get information about specific arguements.",
	Run:   Info,
}

func Info(cmd *cobra.Command, args []string) {
	setLogFormat()
}
