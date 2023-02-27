package tools

import (
	"fmt"

	"dancheg97.ru/dancheg97/gen-tools/templates/services"
	"dancheg97.ru/dancheg97/gen-tools/utils"
)

func GenerateSqlc(repo string, generate bool) {
	utils.WriteFile("sqlc.yaml", SqlcYaml)
	utils.WriteFile("sqlc.sql", SqlcSql)
	utils.WriteFile("migrations/0001_ini.sql", GooseMigrations)
	utils.WriteFile("postgres/postgres.go", fmt.Sprintf(PostgresGo, repo))
	utils.AppendToMakefile(SqlcMakefile)
	utils.AppendToCompose(services.PostgresYml)
}

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
	docker run --rm -it -v ${pwd}:/wd -w /wd dancheg97.ru/dancheg97/gen-tools:latest sqlc generate

`

const PostgresGo = `package postgres

import (
	"context"
	"errors"

	"%s/gen/postgres"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Postgres struct {
	*pgxpool.Pool
	*sqlc.Queries
}

func Get(conn string) (*Postgres, error) {
	config, err := pgxpool.ParseConfig(conn)
	if err != nil {
		return nil, err
	}

	config.ConnConfig.LogLevel = pgx.LogLevelError
	config.ConnConfig.Logger = logrusadapter.NewLogger(logrus.StandardLogger())
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	sqlc := sqlc.New(pool)
	return &Postgres{
		Pool:    pool,
		Queries: sqlc,
	}, nil
}

func (pg *Postgres) WrapTx(ctx context.Context, txFunc func(p *Postgres) error) error {
	tx, err := pg.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	dbtx := pg.Queries.WithTx(tx)
	defer func() {
		rollbackErr := tx.Rollback(ctx)
		if rollbackErr != nil {
			if !errors.Is(rollbackErr, pgx.ErrTxClosed) {
				logrus.Error("error rolling transaction back", rollbackErr)
			}
		}
	}()
	err = txFunc(&Postgres{
		Pool:    pg.Pool,
		Queries: dbtx,
	})
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (pg *Postgres) CheckConstraint(err error, cst string) bool {
	if err == nil {
		return false
	}
	var pgErr *pgconn.PgError
	ok := errors.As(err, &pgErr)
	if ok {
		return pgErr.ConstraintName == cst
	}
	return false
}
`
