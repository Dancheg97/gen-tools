package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-tools",
	Short: "🧰 Tool for generating go project templates.",
	Long: `🧰 Tool for generating go project templates.

	LICENSE

	go-tools  Copyright (C) 2023  Dancheg97
	
	This program comes with ABSOLUTELY NO WARRANTY; for details 'use -h flag'.
	This is free software, and you are welcome to redistribute it
	under certain conditions; watch license in repo for details.

📃 Included tools:

This tool allows to generate prepared go code for interaction with following
infrastructure components (go-tools gen ...):

- drone
- golang-ci
- cobra
- viper
- gRPC
- buf
- docker
- compose
- pgx
- sqlc
- redis
- nats
- license
`,
}

var flags = []Flag{}

func Execute() {
	for _, flag := range flags {
		AddFlag(flag)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
