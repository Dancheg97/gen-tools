package cmd

import (
	"os"

	"dancheg97.ru/dancheg97/gen-tools/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var infraCmd = &cobra.Command{
	Use:   "infr",
	Short: "🐳 Generate all infrastructure in single command",
	Run:   Infr,
}

func init() {
	rootCmd.AddCommand(infraCmd)
}

func Infr(cmd *cobra.Command, args []string) {
	PreventDefault(`mail`, viper.GetString(`mail`), `mail@example.com`)
	PreventDefault(`domain`, viper.GetString(`domain`), `example.com`)
	PreventDefault(`user`, viper.GetString(`user`), `admin`)
	PreventDefault(`pass`, viper.GetString(`pass`), `password`)

	Gen(cmd, []string{
		"readme",
		"service-nginx",
		"service-gitea",
		"service-drone",
		"service-mkdocs",
		"service-kuma",
		"service-dozzle",
	})

	logrus.Info("Obtaining certificates")
	err := utils.SystemCall(`sh certs.sh`)
	utils.CheckErr(err)

	logrus.Info("Generating gitea/mkdocs logo")

	logrus.Info("to run infrastructure run: docker compose up")
}

func PreventDefault(name string, actual string, initial string) {
	if actual == initial {
		logrus.Error(name + " was not applied, value is: " + actual)
		os.Exit(1)
	}
}
