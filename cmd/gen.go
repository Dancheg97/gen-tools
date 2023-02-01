package cmd

import (
	"os"

	"gitea.dancheg97.ru/templates/go-tools/templates"
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
	Long: `📃 Generate template components

📃 Options you can specify under 'gen' command:

This tool allows to generate prepared go code for interaction with following
infrastructure components (go-tools gen [options]):

- cli - incluudes cobra and viper
- drone - includes drone template for CI-CD
- lint - includes golanglint-ci linter for go code
- grpc - includes proto and buf files for generation
- docker - includes 2 stage dockerfile and compose for ease of development
- pg - includes pgx module in porstgres, sqlc for generation and goose for migrations
- redis - includes redis template
- nats - includes consumer and producer nats template
- make - adds makefile with some prepared commands
- gpl - adds GPLv3 license to project
- mit - adds MIT license to project

📒 Recommened to always use in projects:

go-tools gen cli lint docker makefile license
`,
}

func Gen(cmd *cobra.Command, args []string) {
	setLogFormat()

	errs := []error{}

	for _, arg := range args {
		switch arg {
		case "cli":
			errs = append(errs, os.WriteFile("main.go", []byte(templates.CliMainGo), 0600))
			errs = append(errs, os.MkdirAll("cmd", os.ModePerm))
			errs = append(errs, os.WriteFile("cmd/flags.go", []byte(templates.CliFlagsGo), 0600))
			errs = append(errs, os.WriteFile("cmd/run.go", []byte(templates.CliRunGo), 0600))
			errs = append(errs, os.WriteFile("cmd/root.go", []byte(templates.CliRootGo), 0600))
		case "drone":
			errs = append(errs, os.WriteFile(".drone.yml", []byte(templates.DroneYml), 0600))
		case "lint":
			errs = append(errs, os.WriteFile(".golangci.yml", []byte(templates.GolangCiYml), 0600))
		case "grpc":
			errs = append(errs, os.WriteFile("buf.yaml", []byte(templates.BufYaml), 0600))
			errs = append(errs, os.WriteFile("buf.gen.yaml", []byte(templates.BufGenYaml), 0600))
		case "docker":
			errs = append(errs, os.WriteFile("Dockerfile", []byte(templates.Dockerfile), 0600))
			errs = append(errs, os.WriteFile("docker-compose.yml", []byte(templates.DockerCompose), 0600))
		case "pg":
			errs = append(errs, os.WriteFile("sqlc.yaml", []byte(templates.SqlcYaml), 0600))
			errs = append(errs, os.WriteFile("sqlc.sql", []byte(templates.SqlcSql), 0600))
			errs = append(errs, os.MkdirAll("migrations", os.ModePerm))
			errs = append(errs, os.WriteFile("migrations/0001_ini.sql", []byte(templates.MigrationSql), 0600))
			errs = append(errs, os.MkdirAll("postgres", os.ModePerm))
			errs = append(errs, os.WriteFile("postgres/postgres.go", []byte(templates.PostgresGo), 0600))
		case "redis":
			errs = append(errs, os.MkdirAll("redis", os.ModePerm))
			errs = append(errs, os.WriteFile("redis/redis.go", []byte(templates.RedisGo), 0600))
		case "nats":
			errs = append(errs, os.MkdirAll("nats", os.ModePerm))
			errs = append(errs, os.WriteFile("nats/consumer.go", []byte(templates.NatsConsumerGo), 0600))
			errs = append(errs, os.WriteFile("nats/producer.go", []byte(templates.NatsProducerGo), 0600))
		case "make":
			errs = append(errs, os.WriteFile("Makefile", []byte(templates.Makefile), 0600))
		case "gpl":
			errs = append(errs, os.WriteFile("LICENSE", []byte(templates.LicenseGPLv3), 0600))
		case "mit":
			errs = append(errs, os.WriteFile("LICENSE", []byte(templates.LicenseMIT), 0600))
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
