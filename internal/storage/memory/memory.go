package memory

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-memdb"

	"github.com/andranikuz/shortener/internal/models"
)

type MemoryDB struct {
	memory *memdb.MemDB
}

func NewMemoryDB() (*MemoryDB, error) {
	db := MemoryDB{}

	return &db, nil
}

func (db *MemoryDB) Migrate(ctx context.Context) error {
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
				},
			},
		},
	}
	// Create database
	memory, err := memdb.NewMemDB(schema)
	if err != nil {
		return err
	}
	db.memory = memory

	return nil
}

// Save url
func (db *MemoryDB) Save(ctx context.Context, url models.URL) error {
	txn := db.memory.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("url", url); err != nil {
		return fmt.Errorf("saving url error id=%s url=%s", url.ID, url.FullURL)
	}
	txn.Commit()

	return nil
}

// Get url
func (db *MemoryDB) Get(ctx context.Context, id string) (*models.URL, error) {
	txn := db.memory.Txn(false)
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

// Get url by full_url
func (db *MemoryDB) GetByFullURL(ctx context.Context, fullURL string) (*models.URL, error) {
	txn := db.memory.Txn(false)
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

// Save batch of urls
func (db *MemoryDB) SaveBatch(ctx context.Context, urls []models.URL) error {
	for _, url := range urls {
		if err := db.Save(ctx, url); err != nil {
			return err
		}
	}

	return nil
}
