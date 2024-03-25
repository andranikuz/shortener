package storage

import (
	"fmt"
	"github.com/andranikuz/shortener/internal/models"

	"github.com/hashicorp/go-memdb"
)

var db *memdb.MemDB

// Init memory DB
func Init() error {
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
	var err error
	// Create database
	db, err = memdb.NewMemDB(schema)
	if err != nil {
		return fmt.Errorf("init DB error %s", err.Error())
	}

	return nil
}

// Save url
func Save(url models.URL) error {
	txn := db.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("url", url); err != nil {
		return fmt.Errorf("saving url error id=%s url=%s", url.ID, url.FullURL)
	}
	txn.Commit()

	return nil
}

// Get url
func Get(id string) (*models.URL, error) {
	txn := db.Txn(false)
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
