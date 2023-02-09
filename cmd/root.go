package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gen-tools",
	Short: "🧰 Tool for generating project templates written in go.",
	Long: `🧰 CLI Tool for generating different project modules for templating
complex systems with ease.

	LICENSE

	gen-tools  Copyright (C) 2023  Dancheg97
	
	This program comes with ABSOLUTELY NO WARRANTY; for details 'use -h flag'.
	This is free software, and you are welcome to redistribute it
	under certain conditions; watch license in repo for details.

📃 You can get started with following commands:

	gen-tools gen [args] - generate component
	gen-tools go - generate go project
	gen-tools infr - generate compose with infrastructure
	gen-tools list - list possible options for generation

Examples:

gen-tools gen drone gpl

gen-tools go --repo myrepo.com/me/tool

gen-tools infr --name Nice --domain nice.org --user admin --pass SeCReT --email he@he.org

gen-tools infr --name Nice --domain nice.org --user admin --pass SeCReT --email he@he.org

`,
}

var flags = []Flag{
	{
		Cmd:         rootCmd,
		Name:        "name",
		Env:         "NAME",
		Value:       "Project",
		Description: "📜 project name, used in devops overrides",
	},
	{
		Cmd:         rootCmd,
		Name:        "repo",
		Env:         "REPO",
		Value:       "example.com/owner/name",
		Description: "📂 repository for go project, used in links and go mod init",
	},
	{
		Cmd:         rootCmd,
		Name:        "domain",
		Env:         "DOMAIN",
		Value:       "example.com",
		Description: "🌐 web domain that is used to obtain certificates",
	},
	{
		Cmd:         rootCmd,
		Name:        "user",
		Env:         "USER",
		Value:       "admin",
		Description: "🛡️ main admin user that is used for authentication",
	},
	{
		Cmd:         rootCmd,
		Name:        "pass",
		Env:         "PASS",
		Value:       "Admin%1Pass",
		Description: "❔ password for admin account",
	},
	{
		Cmd:         rootCmd,
		Name:        "mail",
		Env:         "MAIL",
		Value:       "mail@example.com",
		Description: "📧 email that is used for acme when obtaining certificates",
	},
	{
		Cmd:         rootCmd,
		Name:        "gitea",
		Env:         "GITEA",
		Value:       "gitea.example.com",
		Description: "🍵 gitea link, used for drone template generating",
	},
}

func Execute() {
	for _, flag := range flags {
		AddFlag(flag)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
