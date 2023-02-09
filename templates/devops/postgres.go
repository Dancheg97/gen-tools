package devops

import "dancheg97.ru/templates/gen-tools/templates/utils"

func GeneratePostgres() {
	utils.AppendToCompose(PostgresYml)
}

const PostgresYml = `  migrator:
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
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: db
    healthcheck:
      test: [ "CMD-SHELL", "pg_isready" ]
      interval: 1s
      timeout: 1s
      retries: 5
    ports:
      - 7002:5432
`
