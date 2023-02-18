package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gen-tools",
	Short: "🧰 Tool for generating project templates written in go.",
	Long: `🧰 CLI Tool for generating different project modules for templating complex systems with ease.

	gen-tools  Copyright (C) 2023  Dancheg97
	
	This program comes with ABSOLUTELY NO WARRANTY; for details 'use -h flag'.
	This is free software, and you are welcome to redistribute it
	under certain conditions; watch license in repo for details.

`,
}

var flags = []Flag{
	{
		Cmd:         rootCmd,
		Name:        "name",
		ShortName:   "n",
		Env:         "NAME",
		Value:       "Project",
		Description: "📜 project name, used in devops overrides",
	},
	{
		Cmd:         rootCmd,
		Name:        "repo",
		ShortName:   "r",
		Env:         "REPO",
		Value:       "example.com/owner/name",
		Description: "📂 repository for go project, used in refs and go mod init",
	},
	{
		Cmd:         rootCmd,
		Name:        "domain",
		ShortName:   "d",
		Env:         "DOMAIN",
		Value:       "example.com",
		Description: "🌐 web domain that is used to obtain certificates",
	},
	{
		Cmd:         rootCmd,
		Name:        "user",
		ShortName:   "u",
		Env:         "USER",
		Value:       "admin",
		Description: "🛡️  main admin user that is used for authentication",
	},
	{
		Cmd:         rootCmd,
		Name:        "pass",
		ShortName:   "p",
		Env:         "PASS",
		Value:       "password",
		Description: "❔ password for admin account",
	},
	{
		Cmd:         rootCmd,
		Name:        "mail",
		ShortName:   "m",
		Env:         "MAIL",
		Value:       "mail@example.com",
		Description: "📧 email that is used for acme when obtaining certificates",
	},
	{
		Cmd:         rootCmd,
		Name:        "gitea",
		ShortName:   "g",
		Env:         "GITEA",
		Value:       "gitea.example.com",
		Description: "🍵 gitea link, used for drone template generating",
	},
	{
		Cmd:         rootCmd,
		Name:        "generate",
		Env:         "GENERATE",
		Value:       "true",
		IsRequired:  false,
		Type:        "bool",
		Description: "📃 wether generation commands (buf/sqlc) would be activated",
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
