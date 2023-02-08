package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"dancheg97.ru/templates/gen-tools/templates"
	"dancheg97.ru/templates/gen-tools/templates/devops"
	"dancheg97.ru/templates/gen-tools/templates/golang"
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
	setLogFormat()

	for _, arg := range args {
		switch arg {

		// OVERALL
		case "drone":
			WriteFile(".drone.yml", templates.DroneYml)
		case "make":
			WriteFile("Makefile", templates.Makefile)
		case "gpl":
			WriteFile("LICENSE", templates.LicenseGPLv3)
		case "mit":
			WriteFile("LICENSE", templates.LicenseMIT)
		case "pkgbuild":
			WriteFile("PKGBUILD", templates.Pkgbuild)

		// DEVOPS
		case "compose-gitea":
			AppendToCompose(devops.GiteaYaml)
			WriteFile(`gitea/gitea/templates/home.tmpl`, devops.GiteaHomeTmpl)
			WriteFile(`gitea/gitea/templates/custom/body_outer_pre.tmpl`, devops.GiteaThemeParkTmpl)
			WriteFile(`gitea/gitea/public/css/theme-earl-grey.css`, devops.GiteaEarlGrayCss)
		case "compose-nginx":
			WriteFile("lego.sh", devops.LegoSh)
			AppendToCompose(devops.NginxYaml)
			WriteFile(`nginx/nginx.conf`, devops.NginxConf)
		case "compose-pacman":
			AppendToCompose(devops.PacmanYaml)
		case "compose-pocketbase":
			AppendToCompose(devops.PocketbaseYaml)
		case "compose-nats":
			AppendToCompose(devops.NatsYaml)
		case "compose-postgres":
			AppendToCompose(devops.PostgresYml)
		case "compose-redis":
			AppendToCompose(devops.RedisYaml)
		case "compose-drone":
			AppendToCompose(devops.DroneYaml)
		case "compose-mkdocs":
			AppendToCompose(devops.MkDocsCompose)
			WriteFile(`mkdocs/mkdocs.yml`, devops.MkDocsConfigYaml)
			WriteFile(`mkdocs/docs/stylesheets/extra.css`, devops.MkDocsCss)

		// GOLANG
		case "go-lint":
			WriteFile(".golangci.yml", golang.GolangCiYml)
		case "go-grpc":
			WriteFile("buf.yaml", golang.BufYaml)
			WriteFile("buf.gen.yaml", golang.BufGenYaml)
			WriteFile("proto/v1/example.proto", golang.GrpcProto)
			AppendToMakefile(golang.BufMake)
			SystemCall("buf generate")
		case "go-docker":
			WriteFile("Dockerfile", golang.Dockerfile)
			WriteFile("docker-compose.yml", golang.DockerCompose)
		case "go-sqlc":
			WriteFile("sqlc.yaml", golang.SqlcYaml)
			WriteFile("sqlc.sql", golang.SqlcSql)
			WriteFile("migrations/0001_ini.sql", golang.GooseMigrations)
			AppendToMakefile(golang.SqlcMakefile)
			SystemCall("sqlc generate")
		case "go-redis":
			WriteFile("redis/redis.go", golang.RedisGo)
		case "go-nats":
			WriteFile("nats/nats.go", fmt.Sprintf(golang.NatsWrapperGo, viper.GetString("repo")))
		case "go-cli":
			WriteFile("main.go", fmt.Sprintf(golang.CliMainGo, viper.GetString("repo")))
			WriteFile("cmd/flags.go", golang.CliFlagsGo)
			WriteFile("cmd/run.go", golang.CliRunGo)
			WriteFile("cmd/root.go", golang.CliRootGo)
			SystemCall("go mod init " + viper.GetString("repo"))
			SystemCall("go mod tidy")

		// UNKNOWN
		default:
			logrus.Error("unknown arguement: ", arg)
		}
	}

	SystemCall(`sudo chmod a+rwx -R .`)

	logrus.Info("template generation finished")
}

func WriteFile(file string, content string) {
	PrepareDir(file)
	err := os.WriteFile(file, []byte(content), 0o600)
	checkErr(err)
	logrus.Info("File generated: ", file)
}

func AppendToFile(file string, content string) {
	PrepareDir(file)
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	checkErr(err)

	_, err = f.WriteString(content)
	checkErr(err)
	logrus.Info("File modified: ", file)
}

func AppendToCompose(content string) {
	const compose = `docker-compose.yml`
	if _, err := os.Stat(compose); errors.Is(err, os.ErrNotExist) {
		WriteFile(compose, "services:\n")
	}
	AppendToFile(compose, content)
}

func AppendToMakefile(content string) {
	const makefile = `Makefile`
	if _, err := os.Stat(makefile); errors.Is(err, os.ErrNotExist) {
		WriteFile(makefile, "pwd := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))\n\n")
	}
	AppendToFile(makefile, content)
}

func PrepareDir(filePath string) {
	if len(strings.Split(filePath, `/`)) != 1 {
		splitted := strings.Split(filePath, `/`)
		path := strings.Join(splitted[0:len(splitted)-1], `/`)
		err := os.MkdirAll(path, os.ModePerm)
		checkErr(err)
	}
}

func SystemCall(cmd string) {
	logrus.Info("Executing system call: ", cmd)
	if os.Getenv("IN_DOCKER") != "true" {
		cmd = "docker run --rm -v $(pwd):/wd -w /wd dancheg97.ru/templates/gen-tools:latest " + cmd
	}
	commad := exec.Command("bash", "-c", cmd)
	commad.Stdout = logrus.StandardLogger().Writer()
	commad.Stderr = logrus.StandardLogger().Writer()
	checkErr(commad.Run())
}
