package tools

import (
	"fmt"

	"dancheg97.ru/dancheg97/gen-tools/utils"
)

func GenerateGoCliTemplate(repo string) {
	utils.WriteFile("main.go", fmt.Sprintf(cliMainGo, repo))
	utils.WriteFile("cmd/flags.go", cliFlagsGo)
	utils.WriteFile("cmd/run.go", cliRunGo)
	utils.WriteFile("cmd/root.go", cliRootGo)
	utils.SystemCall("go mod init " + repo)
	utils.SystemCall("go mod tidy")
}

const cliMainGo = `package main

import "%s/cmd"

func main() {
	cmd.Execute()
}
`

const cliRunGo = `package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "🚀 Run this awesome tool",
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	var (
	// filePort = viper.GetInt("second")
	// repoName = viper.GetString("first")
	)
}
`

const cliRootGo = `package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gen-tools",
	Short: "🧰 Cli tool for something awesome.",
	Long:  "long example",
}

var flags = []Flag{
	{
		Cmd:         rootCmd,
		Name:        "flg",
		ShortName:   "f",
		Env:         "FLG",
		Value:       "value",
		Description: "📄 cool description",
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
`

const cliFlagsGo = `package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Short description of contents for command.
type Flag struct {
	// Cobra command that we will bound our cmd to
	Cmd *cobra.Command
	// Name of command in CLI
	Name string
	// Optional short name for command, leave empty to skip short name
	ShortName string
	// Environment variable to read from
	Env string
	// Regular name for the flag
	Value string
	// Wether this value should be provided by user
	IsRequired bool
	// Leave empty if type is string: ["", "strarr", "bool"]
	Type string
	// Description for flag
	Description string
}

// Function to add new command to CLI tool.
func AddFlag(cmd Flag) {
	if cmd.Type == "" {
		cmd.Cmd.PersistentFlags().StringP(cmd.Name, cmd.ShortName, cmd.Value, cmd.Description)
		err := viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
		checkErr(err)
	}

	if cmd.Type == "strarr" {
		cmd.Cmd.PersistentFlags().StringArrayP(cmd.Name, cmd.ShortName, nil, cmd.Description)
		err := viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
		checkErr(err)
	}

	if cmd.Type == "bool" {
		cmd.Cmd.PersistentFlags().BoolP(cmd.Name, cmd.ShortName, false, cmd.Description)
		err := viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
		checkErr(err)
	}

	if cmd.Type == "int" {
		if cmd.Value != "" {
			i, err := strconv.Atoi(cmd.Value)
			if err != nil {
				err = fmt.Errorf("value for flag "+cmd.Name+" should be int: %w", err)
				checkErr(err)
			}
			cmd.Cmd.PersistentFlags().IntP(cmd.Name, cmd.ShortName, i, cmd.Description)
			err = viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
			checkErr(err)
			return
		}
		cmd.Cmd.PersistentFlags().IntP(cmd.Name, cmd.ShortName, 0, cmd.Description)
		err := viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
		checkErr(err)
	}

	if cmd.Env != "" {
		err := viper.BindEnv(cmd.Name, cmd.Env)
		checkErr(err)
	}

	if cmd.IsRequired {
		err := cmd.Cmd.MarkFlagRequired(cmd.Name)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func setLogFormat() {
	switch viper.GetString("logs-fmt") {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	case "pretty":
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			DisableQuote:  true,
			FullTimestamp: true,
		})
	default:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}
`
