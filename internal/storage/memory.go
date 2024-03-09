package storage

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
)

type URL struct {
	ID      string
	FullURL string
}

var db *memdb.MemDB

// Init memory DB
func Init() {
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
				},
			},
		},
	}
	var err error
	// Create database
	db, err = memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
}

// Save url
func Save(url URL) error {
	txn := db.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("url", url); err != nil {
		return fmt.Errorf("saving url error id=%s url=%s", url.ID, url.FullURL)
	}
	txn.Commit()

	return nil
}

// Get url
func Get(id string) (*URL, error) {
	txn := db.Txn(false)
	defer txn.Abort()
	raw, err := txn.First("url", "id", id)
	if err != nil {
		return nil, fmt.Errorf("getting index %s error", id)
	}

	url, ok := raw.(URL)
	if !ok {
		return nil, fmt.Errorf("index %s not found", id)
	}

	return &url, nil
}
