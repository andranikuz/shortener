package memory

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-memdb"

	"github.com/andranikuz/shortener/internal/models"
)

// MemoryStorage репозиторий в памяти.
type MemoryStorage struct {
	memory *memdb.MemDB
}

// NewMemoryStorage функция инициализации MemoryStorage.
func NewMemoryStorage() (*MemoryStorage, error) {
	db := MemoryStorage{}
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"url": &memdb.TableSchema{
				Name: "url",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"url": &memdb.IndexSchema{
						Name:    "url",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "FullURL"},
					},
					"user_id": &memdb.IndexSchema{
						Name:    "user_id",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "UserID"},
					},
				},
			},
		},
	}
	// Create database
	memory, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}
	db.memory = memory

	return &db, nil
}

// Save метод сохранения URL.
func (storage *MemoryStorage) Save(ctx context.Context, url models.URL) error {
	txn := storage.memory.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("url", url); err != nil {
		return fmt.Errorf("saving url error id=%s url=%s", url.ID, url.FullURL)
	}
	txn.Commit()

	return nil
}

// Get метод получения URL по идентификатору.
func (storage *MemoryStorage) Get(ctx context.Context, id string) (*models.URL, error) {
	txn := storage.memory.Txn(false)
	defer txn.Abort()
	raw, err := txn.First("url", "id", id)
	if err != nil {
		return nil, fmt.Errorf("getting index id=%s error", id)
	}

	url, ok := raw.(models.URL)
	if !ok {
		return nil, fmt.Errorf("index %s not found", id)
	}

	return &url, nil
}

// GetByFullURL метод получения URL по послной ссылке.
func (storage *MemoryStorage) GetByFullURL(ctx context.Context, fullURL string) (*models.URL, error) {
	txn := storage.memory.Txn(false)
	defer txn.Abort()
	raw, err := txn.First("url", "url", fullURL)
	if err != nil {
		return nil, fmt.Errorf("getting index fullURL=%s error", fullURL)
	}

	url, ok := raw.(models.URL)
	if !ok {
		return nil, fmt.Errorf("index %s not found", fullURL)
	}

	return &url, nil
}

// SaveBatch метод сохранения массива URL.
func (storage *MemoryStorage) SaveBatch(ctx context.Context, urls []models.URL) error {
	for _, url := range urls {
		if err := storage.Save(ctx, url); err != nil {
			return err
		}
	}

	return nil
}

// GetByUserID метод полуения списка URL по userID.
func (storage *MemoryStorage) GetByUserID(ctx context.Context, userID string) ([]models.URL, error) {
	txn := storage.memory.Txn(false)
	defer txn.Abort()
	rows, err := txn.Get("url", "user_id", userID)
	if err != nil {
		return nil, fmt.Errorf("getting index userID=%s error", userID)
	}
	var urls []models.URL
	for obj := rows.Next(); obj != nil; obj = rows.Next() {
		url := obj.(models.URL)
		urls = append(urls, url)
	}

	return urls, nil
}

// DeleteURLs метод удаления массива URLs.
func (storage *MemoryStorage) DeleteURLs(ctx context.Context, ids []string, userID string) error {
	return nil
}

// Ping метод проверки статуса соединения.
func (storage MemoryStorage) Ping() error {
	return nil
}
