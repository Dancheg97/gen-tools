package golang

const SqlcSql = `-- name: SelectSomething :one

SELECT * FROM something WHERE id = $1;
`

const SqlcYaml = `version: 1
packages:
  - path: "gen/sqlc"
    name: "sqlc"
    engine: "postgresql"
    schema: "migrations/"
    queries: "sqlc.sql"
    emit_interface: true
    emit_prepared_queries: true
    sql_package: "pgx/v4"
overrides:
  - go_type:
      type: "int32"
      pointer: true
    db_type: "pg_catalog.int4"
    nullable: true
  - go_type:
      type: "float64"
      pointer: true
    db_type: "pg_catalog.float8"
    nullable: true
  - go_type:
      type: "float32"
      pointer: true
    db_type: "pg_catalog.float4"
    nullable: true
  - go_type:
      type: "string"
      pointer: true
    db_type: "pg_catalog.varchar"
    nullable: true
  - go_type:
      type: "string"
      pointer: true
    db_type: "text"
    nullable: true
`

const PostgresGo = `package postgres

import (
	"context"
	"errors"

	"gitea.dancheg97.ru/dancheg97/go-sqlc/gen/sqlc"
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

const MigrationSql = `-- +goose Up
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
