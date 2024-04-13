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

func (db *PostgresDB) Migrate(ctx context.Context) error {
	if _, err := db.DB.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS public.url (
			id varchar NOT NULL,
			full_url varchar NOT NULL
		)
	`); err != nil {
		return err
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

// Save batch of urls
func (db *PostgresDB) SaveBatch(ctx context.Context, urls []models.URL) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO url (id, full_url) VALUES ($1,$2)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, url := range urls {
		_, err = stmt.ExecContext(ctx, url.ID, url.FullURL)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
