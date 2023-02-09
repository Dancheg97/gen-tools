package cmd

import (
	"dancheg97.ru/templates/gen-tools/templates"
	"dancheg97.ru/templates/gen-tools/templates/devops"
	"dancheg97.ru/templates/gen-tools/templates/gogen"
	"dancheg97.ru/templates/gen-tools/templates/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "📃 Generate template components",
	Run:   Gen,
}

func init() {
	rootCmd.AddCommand(genCmd)
}

func Gen(cmd *cobra.Command, args []string) {
	repo := viper.GetString(`repo`)
	mail := viper.GetString(`mail`)
	domain := viper.GetString(`domain`)
	user := viper.GetString(`user`)
	pass := viper.GetString(`pass`)
	gitea := viper.GetString(`gitea`)

	setLogFormat()

	for _, arg := range args {
		processArguement(repo, mail, domain, user, pass, gitea, arg)
	}

	utils.SystemCall(`sudo chmod a+rwx -R .`)

	logrus.Info("template generation finished")
}

func processArguement(repo, mail, domain, user, pass, gitea, arg string) {
	switch arg {
	case "drone":
		templates.GenerateDroneYml(gitea)
	case "make":
		utils.WriteFile("Makefile", templates.Makefile)
	case "gpl":
		utils.WriteFile("LICENSE", templates.LicenseGPLv3)
	case "mit":
		utils.WriteFile("LICENSE", templates.LicenseMIT)
	case "pkgbuild":
		utils.WriteFile("PKGBUILD", templates.Pkgbuild)
	case "compose-gitea":
		devops.GenerateGitea(mail, domain)
	case "compose-nginx":
		devops.GenerateNginx()
	case "compose-pacman":
		devops.GeneratePacman(mail, domain)
	case "compose-pocketbase":
		devops.GeneratePocketbase(mail, domain)
	case "compose-nats":
		devops.GenerateNats()
	case "compose-postgres":
		devops.GeneratePostgres(user, pass)
	case "compose-redis":
		devops.GenerateRedis()
	case "compose-drone":
		devops.GenerateDrone(mail, domain)
	case "compose-mkdocs":
		devops.GenerateMkdocs(mail, domain)
	case "compose-kuma":
		devops.GenerateUptimeKuma(mail, domain)
	case "go-lint":
		gogen.GenerateGolangCi()
	case "go-grpc":
		gogen.GenerateBuf()
	case "go-docker":
		gogen.GenerateGoDocker(repo)
	case "go-sqlc":
		gogen.GenerateSqlc(repo)
	case "go-redis":
		gogen.GenerateRedis()
	case "go-nats":
		gogen.GenerateNats(repo)
	case "go-cli":
		gogen.GenerateGoCliTemplate(repo)
	default:
		logrus.Error("unknown arguement: ", arg)
	}
}
