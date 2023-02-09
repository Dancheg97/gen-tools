package gogen

import (
	"fmt"

	"dancheg97.ru/templates/gen-tools/templates/utils"
)

func GenerateGoDocker(name string) {
	utils.WriteFile("Dockerfile", fmt.Sprintf(Dockerfile, name, name, name))
	utils.WriteFile("docker-compose.yml", fmt.Sprintf(DockerCompose, name))
}

const Dockerfile = `FROM golang:1.19.1 AS build
WORKDIR /
COPY . .
RUN go build -a -o %s .
FROM alpine:3.16.2
COPY --from=build /%s .
CMD /%s`

const DockerCompose = `version: '3.9'

services:
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
