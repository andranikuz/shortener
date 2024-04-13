package postgres

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/models"
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

func (db *PostgresDB) Migrate() error {
	var tableExist bool
	row := db.DB.QueryRow(`
		SELECT EXISTS (
    		SELECT 1 FROM information_schema.tables 
    		WHERE table_name = 'url'
		) AS table_exists
	`)
	err := row.Scan(&tableExist)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	if !tableExist {
		if _, err = db.DB.Exec(`
			CREATE TABLE public.url (
				id varchar NOT NULL,
				full_url varchar NOT NULL
			)
		`); err != nil {
			return err
		}
	}

	return nil
}

// Save url
func (db *PostgresDB) Save(ctx context.Context, url models.URL) error {
	if _, err := db.DB.ExecContext(ctx, `
			INSERT INTO url (id, full_url)
			VALUES ($1, $2)
		`, url.ID, url.FullURL); err != nil {
		return err
	}

	return nil
}

// Get url
func (db *PostgresDB) Get(ctx context.Context, id string) (*models.URL, error) {
	row := db.DB.QueryRowContext(ctx, `SELECT id, full_url FROM url where id = $1`, id)
	// готовим переменную для чтения результата
	var url models.URL
	err := row.Scan(&url.ID, &url.FullURL) // разбираем результат
	if err != nil {
		panic(err)
	}

	return &url, nil
}
