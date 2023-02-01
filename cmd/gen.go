package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "📃 Generate template components",
	Run:   Gen,
	Long: `📃 Generate template components

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

func Gen(cmd *cobra.Command, args []string) {
	setLogFormat()
}
