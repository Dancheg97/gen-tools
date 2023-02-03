package cmd

import (
	"os"

	"gitea.dancheg97.ru/templates/gen-tools/templates"
	"gitea.dancheg97.ru/templates/gen-tools/templates/arch"
	"gitea.dancheg97.ru/templates/gen-tools/templates/golang"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:     "gen",
	Short:   "📃 Generate template components",
	Run:     Gen,
	Example: "drone ",
}

func Gen(cmd *cobra.Command, args []string) {
	setLogFormat()

	errs := []error{}

	for _, arg := range args {
		switch arg {
		// OVERALL
		case "drone":
			errs = append(errs, os.WriteFile(".drone.yml", []byte(templates.DroneYml), 0600))
		case "make":
			errs = append(errs, os.WriteFile("Makefile", []byte(templates.Makefile), 0600))
		case "gpl":
			errs = append(errs, os.WriteFile("LICENSE", []byte(templates.LicenseGPLv3), 0600))
		case "mit":
			errs = append(errs, os.WriteFile("LICENSE", []byte(templates.LicenseMIT), 0600))
		// GOLANG
		case "go-cli":
			errs = append(errs, os.WriteFile("main.go", []byte(golang.CliMainGo), 0600))
			errs = append(errs, os.MkdirAll("cmd", os.ModePerm))
			errs = append(errs, os.WriteFile("cmd/flags.go", []byte(golang.CliFlagsGo), 0600))
			errs = append(errs, os.WriteFile("cmd/run.go", []byte(golang.CliRunGo), 0600))
			errs = append(errs, os.WriteFile("cmd/root.go", []byte(golang.CliRootGo), 0600))
		case "go-lint":
			errs = append(errs, os.WriteFile(".golangci.yml", []byte(golang.GolangCiYml), 0600))
		case "go-grpc":
			errs = append(errs, os.WriteFile("buf.yaml", []byte(golang.BufYaml), 0600))
			errs = append(errs, os.WriteFile("buf.gen.yaml", []byte(golang.BufGenYaml), 0600))
		case "go-docker":
			errs = append(errs, os.WriteFile("Dockerfile", []byte(golang.Dockerfile), 0600))
			errs = append(errs, os.WriteFile("docker-compose.yml", []byte(golang.DockerCompose), 0600))
		case "go-pg":
			errs = append(errs, os.WriteFile("sqlc.yaml", []byte(golang.SqlcYaml), 0600))
			errs = append(errs, os.WriteFile("sqlc.sql", []byte(golang.SqlcSql), 0600))
			errs = append(errs, os.MkdirAll("migrations", os.ModePerm))
			errs = append(errs, os.WriteFile("migrations/0001_ini.sql", []byte(golang.MigrationSql), 0600))
			errs = append(errs, os.MkdirAll("postgres", os.ModePerm))
			errs = append(errs, os.WriteFile("postgres/postgres.go", []byte(golang.PostgresGo), 0600))
		case "go-redis":
			errs = append(errs, os.MkdirAll("redis", os.ModePerm))
			errs = append(errs, os.WriteFile("redis/redis.go", []byte(golang.RedisGo), 0600))
		case "go-nats":
			errs = append(errs, os.MkdirAll("nats", os.ModePerm))
			errs = append(errs, os.WriteFile("nats/consumer.go", []byte(golang.NatsConsumerGo), 0600))
			errs = append(errs, os.WriteFile("nats/producer.go", []byte(golang.NatsProducerGo), 0600))
			// CTRL
		case "pkgbuild":
			errs = append(errs, os.WriteFile("PKGBUILD", []byte(arch.Pkgbuild), 0600))
		// UNKNOWN
		default:
			logrus.Error("unknown arguement: ", arg)
		}
	}

	for _, err := range errs {
		if err != nil {
			logrus.Error(err)
		}
	}
	logrus.Info("template generation finished")
}
