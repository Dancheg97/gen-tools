package golang

import (
	"fmt"
	"strings"

	"dancheg97.ru/dancheg97/gen-tools/utils"
)

func GenerateGoDocker(name string) {
	splitted := strings.Split(name, `/`)
	name = splitted[len(splitted)-1]
	utils.WriteFile("Dockerfile", fmt.Sprintf(Dockerfile, name))
	utils.AppendToCompose(fmt.Sprintf(DockerCompose, name))
}

const Dockerfile = `FROM golang:1.19.1
WORKDIR /
COPY . .
RUN go install .
ENTRYPOINT ["%s"]`

const DockerCompose = `
  %s:
    build:
      context: .
    command: run
    environment:
      LOGS_FMT: text
    volumes:
      - ./host:/in/container
    ports:
      - 9080:9080
      - 8080:8080

`
