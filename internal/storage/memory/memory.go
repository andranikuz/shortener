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
		return nil, fmt.Errorf("init DB error %s", err.Error())
	}

	db := MemoryDB{memory}

	return &db, nil
}

func (db *MemoryDB) Migrate() error {
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
