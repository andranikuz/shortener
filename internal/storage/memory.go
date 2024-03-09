package storage

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
)

type Url struct {
	Id  string
	Url string
}

var Database *memdb.MemDB

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
						Indexer: &memdb.StringFieldIndex{Field: "Id"},
					},
				},
			},
		},
	}
	var err error
	// Create database
	Database, err = memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
}

// Save url
func Save(url Url) error {
	txn := Database.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("url", url); err != nil {
		return fmt.Errorf("Saving url error id=%s url=%s", url.Id, url.Url)
	}
	txn.Commit()

	return nil
}

// Get url
func Get(id string) (*Url, error) {
	txn := Database.Txn(false)
	defer txn.Abort()
	raw, err := txn.First("url", "id", id)
	if err != nil {
		return nil, fmt.Errorf("Getting index %s error", id)
	}

	url, ok := raw.(Url)
	if !ok {
		return nil, fmt.Errorf("Index %s not found", id)
	}

	return &url, nil
}
