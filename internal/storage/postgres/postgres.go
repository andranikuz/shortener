package postgres

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/models"
)

// PostgresStorage Postgres репозиторий.
type PostgresStorage struct {
	DB *sql.DB
}

// NewPostgresStorage функция инициализации PostgresStorage.
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

// Migrate метод для создания таблицы url в БД.
func (storage *PostgresStorage) Migrate() error {
	if _, err := storage.DB.Exec(`
		CREATE TABLE IF NOT EXISTS public.url (
			id varchar NOT NULL,
			full_url varchar NOT NULL,
			user_id varchar NULL,
			is_deleted bool NOT NULL
		);
		CREATE UNIQUE INDEX IF NOT EXISTS url_full_url_idx ON public.url USING btree (full_url);
	`); err != nil {
		return err
	}

	return nil
}

// Save метод сохранения URL.
func (storage *PostgresStorage) Save(ctx context.Context, url models.URL) error {
	if _, err := storage.DB.ExecContext(ctx, `
			INSERT INTO url (id, full_url, user_id, is_deleted)
			VALUES ($1, $2, $3, $4)
		`, url.ID, url.FullURL, url.UserID, false); err != nil {
		return err
	}

	return nil
}

// Get метод получения URL по идентификатору.
func (storage *PostgresStorage) Get(ctx context.Context, id string) (*models.URL, error) {
	row := storage.DB.QueryRowContext(ctx, `SELECT id, full_url, is_deleted FROM url where id = $1`, id)
	var url models.URL
	err := row.Scan(&url.ID, &url.FullURL, &url.DeletedFlag)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

// GetByFullURL метод получения URL по послной ссылке.
func (storage *PostgresStorage) GetByFullURL(ctx context.Context, fullURL string) (*models.URL, error) {
	row := storage.DB.QueryRowContext(ctx, `SELECT id, full_url, is_deleted FROM url where full_url = $1`, fullURL)
	var url models.URL
	err := row.Scan(&url.ID, &url.FullURL, &url.DeletedFlag)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

// SaveBatch метод сохранения массива URL.
func (storage *PostgresStorage) SaveBatch(ctx context.Context, urls []models.URL) error {
	tx, err := storage.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO url (id, full_url, user_id, is_deleted) VALUES ($1,$2,$3,$4) ON CONFLICT DO NOTHING`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, url := range urls {
		_, err = stmt.ExecContext(ctx, url.ID, url.FullURL, url.UserID, false)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// GetByUserID метод полуения списка URL по userID.
func (storage *PostgresStorage) GetByUserID(ctx context.Context, userID string) ([]models.URL, error) {
	var urls []models.URL
	rows, err := storage.DB.QueryContext(ctx, `SELECT id, full_url, is_deleted FROM url WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		var url models.URL
		if err = rows.Scan(&url.ID, &url.FullURL, &url.DeletedFlag); err != nil {
			return urls, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

// DeleteURLs метод удаления массива URLs.
func (storage *PostgresStorage) DeleteURLs(ctx context.Context, ids []string, userID string) error {
	tx, err := storage.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		UPDATE url 
		SET is_deleted = true 
		WHERE id = $1
		AND user_id = $2
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, id := range ids {
		_, err = stmt.ExecContext(ctx, id, userID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// Ping метод проверки статуса соединения.
func (storage PostgresStorage) Ping() error {
	return storage.DB.Ping()
}

// GetUsersCount метод получения количества пользователей.
func (storage PostgresStorage) GetUsersCount(ctx context.Context) (int, error) {
	row := storage.DB.QueryRowContext(ctx, `SELECT COUNT(DISTINCT(user_id)) FROM url `)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetURLsCount метод получения количества записей.
func (storage PostgresStorage) GetURLsCount(ctx context.Context) (int, error) {
	row := storage.DB.QueryRowContext(ctx, `SELECT COUNT(user_id) FROM url`)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
