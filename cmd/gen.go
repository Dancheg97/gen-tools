package cmd

import (
	"dancheg97.ru/dancheg97/gen-tools/templates"
	"dancheg97.ru/dancheg97/gen-tools/templates/golang"
	"dancheg97.ru/dancheg97/gen-tools/templates/licenses"
	"dancheg97.ru/dancheg97/gen-tools/templates/services"
	"dancheg97.ru/dancheg97/gen-tools/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	genCmd = &cobra.Command{
		Use:   "gen",
		Short: "🧰 Generate template components",
		Run:   Gen,
	}

	repo   string
	mail   string
	domain string
	user   string
	pass   string
	gitea  string
	logo   string
)

func init() {
	rootCmd.AddCommand(genCmd)
}

func Gen(cmd *cobra.Command, args []string) {
	setLogFormat()

	repo = viper.GetString(`repo`)
	mail = viper.GetString(`mail`)
	domain = viper.GetString(`domain`)
	user = viper.GetString(`user`)
	pass = viper.GetString(`pass`)
	gitea = viper.GetString(`gitea`)
	logo = viper.GetString(`logo`)

	for _, arg := range args {
		processArguement(arg)
	}

	logrus.Info("template generation finished")
}

func processArguement(arg string) {
	switch arg {
	case "readme":
		utils.WriteFile(`README.md`, templates.Readme)
	case "drone":
		templates.GenerateDroneYml(gitea)
	case "license-gpl":
		utils.WriteFile("LICENSE", licenses.GPLv3)
	case "license-mit":
		utils.WriteFile("LICENSE", licenses.MIT)
	case "license-apache":
		utils.WriteFile("LICENSE", licenses.Apache)
	case "pkgbuild":
		utils.WriteFile("PKGBUILD", templates.Pkgbuild)
	case "service-gitea":
		services.GenerateGitea(mail, domain, logo)
	case "service-nginx":
		services.GenerateNginx()
	case "service-pacman":
		services.GeneratePacman(mail, domain)
	case "service-pocketbase":
		services.GeneratePocketbase(mail, domain)
	case "service-nats":
		services.GenerateNats()
	case "service-postgres":
		services.GeneratePostgres(user, pass)
	case "service-redis":
		services.GenerateRedis()
	case "service-drone":
		services.GenerateDrone(mail, domain)
	case "service-mkdocs":
		services.GenerateMkdocs(mail, domain)
	case "service-kuma":
		services.GenerateUptimeKuma(mail, domain)
	case "service-dozzle":
		services.GenerateDozzle(mail, domain, user, pass)
	case "go-lint":
		golang.GenerateGolangCi()
	case "go-grpc":
		golang.GenerateBuf()
	case "go-docker":
		golang.GenerateGoDocker(repo)
	case "go-sqlc":
		golang.GenerateSqlc(repo)
	case "go-redis":
		golang.GenerateRedis()
	case "go-nats":
		golang.GenerateNats(repo)
	case "go-cli":
		golang.GenerateGoCliTemplate(repo)
	default:
		logrus.Error("unknown arguement: ", arg)
	}
}
