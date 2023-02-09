package cmd

import (
	"dancheg97.ru/templates/gen-tools/templates/gogen"
	"dancheg97.ru/templates/gen-tools/templates/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var genCmd = &cobra.Command{
	Use:     "gen",
	Short:   "📃 Generate template components",
	Run:     Gen,
	Example: "gen-tools --repo testgen go-lint go-grpc go-docker go-sqlc go-redis go-cli go-nats",
}

func init() {
	rootCmd.AddCommand(genCmd)
}

func Gen(cmd *cobra.Command, args []string) {
	var (
		repo = viper.GetString(`repo`)
	)

	setLogFormat()

	for _, arg := range args {
		switch arg {
		// OVERALL
		// case "drone":
		// 	WriteFile(".drone.yml", templates.DroneYml)
		// case "make":
		// 	WriteFile("Makefile", templates.Makefile)
		// case "gpl":
		// 	WriteFile("LICENSE", templates.LicenseGPLv3)
		// case "mit":
		// 	WriteFile("LICENSE", templates.LicenseMIT)
		// case "pkgbuild":
		// 	WriteFile("PKGBUILD", templates.Pkgbuild)

		// // DEVOPS
		// case "compose-gitea":
		// 	AppendToCompose(devops.GiteaYaml)
		// 	WriteFile(`gitea/gitea/templates/home.tmpl`, devops.GiteaHomeTmpl)
		// 	WriteFile(`gitea/gitea/templates/custom/body_outer_pre.tmpl`, devops.GiteaThemeParkTmpl)
		// 	WriteFile(`gitea/gitea/public/css/theme-earl-grey.css`, devops.GiteaEarlGrayCss)
		// case "compose-nginx":
		// 	WriteFile("lego.sh", devops.LegoSh)
		// 	AppendToCompose(devops.NginxYaml)
		// 	WriteFile(`nginx/nginx.conf`, devops.NginxConf)
		// case "compose-pacman":
		// 	AppendToCompose(devops.PacmanYaml)
		// case "compose-pocketbase":
		// 	AppendToCompose(devops.PocketbaseYaml)
		// case "compose-nats":
		// 	AppendToCompose(devops.NatsYaml)
		// case "compose-postgres":
		// 	AppendToCompose(devops.PostgresYml)
		// case "compose-redis":
		// 	AppendToCompose(devops.RedisYaml)
		// case "compose-drone":
		// 	AppendToCompose(devops.DroneYaml)
		// case "compose-mkdocs":
		// 	AppendToCompose(devops.MkDocsCompose)
		// 	WriteFile(`mkdocs/mkdocs.yml`, devops.MkDocsConfigYaml)
		// 	WriteFile(`mkdocs/docs/stylesheets/extra.css`, devops.MkDocsCss)
		// GOLANG
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

	utils.SystemCall(`sudo chmod a+rwx -R .`)

	logrus.Info("template generation finished")
}
