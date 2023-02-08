package golang

const SqlcSql = `-- name: SelectSomething :one

SELECT * FROM users WHERE id = $1;
`

const SqlcYaml = `version: "2"
sql:
  - schema: "migrations/"
    queries: "sqlc.sql"
    engine: "postgresql"
    gen:
      go: 
        package: "postgres"
        out: "gen/postgres"
        emit_interface: true
`

const GooseMigrations = `-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    age INTEGER NOT NULL,
    some_id INTEGER NOT NULL
);
-- +goose Down
DROP TABLE user;
`

const SqlcMakefile = `
sqlc:
	docker run --rm -it -v ${pwd}:/wd -w /wd dancheg97.ru/templates/gen-tools:latest sqlc generate

`
