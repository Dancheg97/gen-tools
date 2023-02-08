package golang

const Dockerfile = `FROM golang:1.19.1 AS build
WORKDIR /
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app ./cmd/server/main.go
FROM alpine:3.16.2
COPY --from=build /app .
CMD /app`

const DockerCompose = `version: '3.9'

services:
  app:
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
