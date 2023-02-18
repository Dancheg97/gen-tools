package services

import (
	"fmt"

	"dancheg97.ru/dancheg97/gen-tools/utils"
)

func GeneratePostgres(name string, pass string) {
	utils.AppendToCompose(fmt.Sprintf(PostgresYml, name, pass))
}

const PostgresYml = `
  migrator:
    image: dangdancheg/goose:0.0.1
    volumes:
      - ./migrations:/migrations
    entrypoint:
      [
        "goose",
        "-dir=/migrations",
        "postgres",
        "user=user host=postgres password=password dbname=db sslmode=disable",
        "up"
      ]
    depends_on:
      postgres:
        condition: service_healthy

  postgres:
    image: postgres:14
    environment:
      POSTGRES_USER: %s
      POSTGRES_PASSWORD: %s
      POSTGRES_DB: db
    healthcheck:
      test: [ "CMD-SHELL", "pg_isready" ]
      interval: 1s
      timeout: 1s
      retries: 5
    ports:
      - 7002:5432

`
