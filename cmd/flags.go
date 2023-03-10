package cmd

import (
	"fmt"
	"strconv"

	"dancheg97.ru/dancheg97/gen-tools/utils"
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
		utils.CheckErr(err)
	}

	if cmd.Type == "strarr" {
		cmd.Cmd.PersistentFlags().StringArrayP(cmd.Name, cmd.ShortName, nil, cmd.Description)
		err := viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
		utils.CheckErr(err)
	}

	if cmd.Type == "bool" {
		cmd.Cmd.PersistentFlags().BoolP(cmd.Name, cmd.ShortName, false, cmd.Description)
		err := viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
		utils.CheckErr(err)
	}

	if cmd.Type == "int" {
		if cmd.Value != "" {
			i, err := strconv.Atoi(cmd.Value)
			if err != nil {
				err = fmt.Errorf("value for flag "+cmd.Name+" should be int: %w", err)
				utils.CheckErr(err)
			}
			cmd.Cmd.PersistentFlags().IntP(cmd.Name, cmd.ShortName, i, cmd.Description)
			err = viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
			utils.CheckErr(err)
			return
		}
		cmd.Cmd.PersistentFlags().IntP(cmd.Name, cmd.ShortName, 0, cmd.Description)
		err := viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
		utils.CheckErr(err)
	}

	if cmd.Env != "" {
		err := viper.BindEnv(cmd.Name, cmd.Env)
		utils.CheckErr(err)
	}

	if cmd.IsRequired {
		err := cmd.Cmd.MarkFlagRequired(cmd.Name)
		utils.CheckErr(err)
	}
}

func setLogFormat() {
	switch viper.GetString("logs-fmt") {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors:    true,
			DisableTimestamp: true,
			DisableQuote:     true,
		})
	case "colored":
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:      true,
			DisableQuote:     true,
			DisableTimestamp: true,
		})
	}
}
