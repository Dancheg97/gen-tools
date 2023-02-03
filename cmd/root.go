package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gen-tools",
	Short: "🧰 Tool for generating go project templates.",
	Long: `🧰 Tool for generating go project templates.

	LICENSE

	gen-tools  Copyright (C) 2023  Dancheg97
	
	This program comes with ABSOLUTELY NO WARRANTY; for details 'use -h flag'.
	This is free software, and you are welcome to redistribute it
	under certain conditions; watch license in repo for details.

📃 Options you can specify under 'gen' command:

This tool allows to generate prepared go code for interaction with following
infrastructure components (gen-tools gen [options]):

- cli - incluudes cobra and viper
- drone - includes drone template for CI-CD
- lint - includes golanglint-ci linter for go code
- grpc - includes proto and buf files for generation
- docker - includes 2 stage dockerfile and compose for ease of development
- pg - includes pgx module in porstgres, sqlc for generation and goose for migrations
- redis - includes redis template
- nats - includes consumer and producer nats template
- license - adds GPLv3 license to project
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
