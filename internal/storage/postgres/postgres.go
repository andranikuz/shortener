package postgres

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/models"
)

type PostgresStorage struct {
	DB *sql.DB
}

func NewPostgresStorage(dsn string) (*PostgresStorage, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	postgresDB := PostgresStorage{
		DB: db,
	}

	return &postgresDB, nil
}

func (storage *PostgresStorage) Migrate() error {
	if _, err := storage.DB.Exec(`
		CREATE TABLE IF NOT EXISTS public.url (
			id varchar NOT NULL,
			full_url varchar NOT NULL,
			user_id varchar NULL
		);
		CREATE UNIQUE INDEX IF NOT EXISTS url_full_url_idx ON public.url USING btree (full_url);
	`); err != nil {
		return err
	}

	return nil
}

// Save url
func (storage *PostgresStorage) Save(ctx context.Context, url models.URL) error {
	if _, err := storage.DB.ExecContext(ctx, `
			INSERT INTO url (id, full_url, user_id)
			VALUES ($1, $2, $3)
		`, url.ID, url.FullURL, url.UserID); err != nil {
		return err
	}

	return nil
}

// Get url
func (storage *PostgresStorage) Get(ctx context.Context, id string) (*models.URL, error) {
	row := storage.DB.QueryRowContext(ctx, `SELECT id, full_url FROM url where id = $1`, id)
	var url models.URL
	err := row.Scan(&url.ID, &url.FullURL)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

// Get url by full_url
func (storage *PostgresStorage) GetByFullURL(ctx context.Context, fullURL string) (*models.URL, error) {
	row := storage.DB.QueryRowContext(ctx, `SELECT id, full_url FROM url where full_url = $1`, fullURL)
	var url models.URL
	err := row.Scan(&url.ID, &url.FullURL)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

// Save batch of urls
func (storage *PostgresStorage) SaveBatch(ctx context.Context, urls []models.URL) error {
	tx, err := storage.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO url (id, full_url, user_id) VALUES ($1,$2,$3) ON CONFLICT DO NOTHING`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, url := range urls {
		_, err = stmt.ExecContext(ctx, url.ID, url.FullURL, url.UserID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (storage *PostgresStorage) GetByUserID(ctx context.Context, userID string) ([]models.URL, error) {
	var urls []models.URL
	rows, err := storage.DB.QueryContext(ctx, `SELECT id, full_url FROM url WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		var url models.URL
		if err = rows.Scan(&url.ID, &url.FullURL); err != nil {
			return urls, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}
