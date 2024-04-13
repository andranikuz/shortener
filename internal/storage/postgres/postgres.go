package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(dsn string) (*PostgresDB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	postgresDB := PostgresDB{
		DB: db,
	}

	return &postgresDB, nil
}
