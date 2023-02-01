package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "🤖 Generate template components",
	Run:   Info,
	Long: `🤖 Generate template components

This tool allows to generate prepared go code for interaction with following
infrastructure components (go-tools gen cli drone lint nats):

- cobra&viper
- drone
- golang-ci lint
- gRPC&buf
- docker&compose
- pgx&sqlc
- redis
- nats`,
}

func Info(cmd *cobra.Command, args []string) {
	setLogFormat()
}
