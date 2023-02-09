package cmd

import "github.com/spf13/cobra"

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "🐋 Generate all go related components",
	Run:   Go,
}

func init() {
	rootCmd.AddCommand(goCmd)
}

func Go(cmd *cobra.Command, args []string) {
	Gen(cmd, []string{
		"go-lint",
		"go-grpc",
		"go-docker",
		"go-sqlc",
		"go-redis",
		"go-nats",
		"go-cli",
		"compose-redis",
		"compose-nats",
		"compose-postgres",
	})
}
