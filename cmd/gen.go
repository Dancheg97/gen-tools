package cmd

import (
	"os"

	"gitea.dancheg97.ru/templates/go-tools/templates"
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
- license - adds GPLv3 license to project
`,
}

func Gen(cmd *cobra.Command, args []string) {
	setLogFormat()

	genCli()
}

func genCli() {
	err := os.WriteFile("main.go", []byte(templates.CliMainGo), 0600)
	checkErr(err)
}
