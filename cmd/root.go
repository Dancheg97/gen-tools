package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gen-tools",
	Short: "🧰 Tool for generating go project templates.",
	Long: `🧰 CLI Tool for generating different project modules for templating
complex systems with ease.

	LICENSE

	gen-tools  Copyright (C) 2023  Dancheg97
	
	This program comes with ABSOLUTELY NO WARRANTY; for details 'use -h flag'.
	This is free software, and you are welcome to redistribute it
	under certain conditions; watch license in repo for details.

📃 Options you can specify under 'gen' command:

This tool allows to generate prepared go code for interaction with following
infrastructure components (gen-tools gen [options]):

- drone - includes drone template for CI-CD
- gpl - adds GPLv3 license to project
- mit - adds MIT license to project
- make - adds Makefile to project
- pkgbuild - arch format PKGBUILD for packaging
- go-cli - includes cobra and viper
- go-lint - includes golanglint-ci linter for go code
- go-grpc - includes proto and buf files for generation
- go-docker - includes 2 stage Dockerfile and compose for ease of development
- go-pg - includes pgx module in porstgres, sqlc for generation and goose for migrations
- go-redis - includes redis template
- go-nats - includes consumer and producer nats template

Example:

gen-tools gen drone make gpl

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
