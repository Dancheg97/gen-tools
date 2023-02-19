package cmd

import (
	"dancheg97.ru/dancheg97/gen-tools/templates"
	"dancheg97.ru/dancheg97/gen-tools/templates/services"
	"dancheg97.ru/dancheg97/gen-tools/templates/tools"
	"dancheg97.ru/dancheg97/gen-tools/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "🧰 Generate template components",
	Run:   Gen,
}

var (
	repo   string
	mail   string
	domain string
	user   string
	pass   string
	gitea  string
	logo   string
	gen    bool
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
	gen = viper.GetBool(`generate`)
	logo = viper.GetString(`logo`)

	for _, arg := range args {
		processArguement(arg)
	}

	logrus.Info("template generation finished")
}

func processArguement(arg string) {
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
		services.GenerateGitea(mail, domain, logo)
	case "compose-nginx":
		services.GenerateNginx()
	case "compose-pacman":
		services.GeneratePacman(mail, domain)
	case "compose-pocketbase":
		services.GeneratePocketbase(mail, domain)
	case "compose-nats":
		services.GenerateNats()
	case "compose-postgres":
		services.GeneratePostgres(user, pass)
	case "compose-redis":
		services.GenerateRedis()
	case "compose-drone":
		services.GenerateDrone(mail, domain)
	case "compose-mkdocs":
		services.GenerateMkdocs(mail, domain)
	case "compose-kuma":
		services.GenerateUptimeKuma(mail, domain)
	case "compose-dozzle":
		services.GenerateDozzle(mail, domain, user, pass)
	case "go-lint":
		tools.GenerateGolangCi()
	case "go-grpc":
		tools.GenerateBuf(gen)
	case "go-docker":
		tools.GenerateGoDocker(repo)
	case "go-sqlc":
		tools.GenerateSqlc(repo, gen)
	case "go-redis":
		tools.GenerateRedis()
	case "go-nats":
		tools.GenerateNats(repo)
	case "go-cli":
		tools.GenerateGoCliTemplate(repo)
	default:
		logrus.Error("unknown arguement: ", arg)
	}
}
