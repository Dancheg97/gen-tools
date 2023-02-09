package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var infraCmd = &cobra.Command{
	Use:   "infr",
	Short: "🎚️ Generate all infrastructure in single command",
	Run:   Infr,
}

func init() {
	rootCmd.AddCommand(infraCmd)
}

func Infr(cmd *cobra.Command, args []string) {
	Gen(cmd, []string{
		"compose-nginx",
		"compose-gitea",
		"compose-pocketbase",
		"compose-drone",
		"compose-mkdocs",
		"compose-pacman",
		"compose-kuma",
	})
	logrus.Info("to obtain certificates run: sh certs.sh")
	logrus.Info("to run infrastructure run: docker compose up")
}
