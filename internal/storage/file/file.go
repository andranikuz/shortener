package file

import (
	"context"
	"encoding/json"
	"os"

	"github.com/andranikuz/shortener/internal/models"
)

type FileDB struct {
	filePath string
}

func NewFileDB(path string) (*FileDB, error) {
	db := FileDB{path}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	file.Close()

	return &db, nil
}

func (db *FileDB) Migrate() error {
	return nil
}

// Save url
func (db *FileDB) Save(ctx context.Context, url models.URL) error {
	data, err := json.Marshal(&url)
	if err != nil {
		return err
	}
	p, err := newProducer(db.filePath)
	if err != nil {
		return err
	}
	if err := p.write(data); err != nil {
		return err
	}

	return nil
}

// Get url
func (db *FileDB) Get(ctx context.Context, id string) (*models.URL, error) {
	c, err := newConsumer(db.filePath)
	if err != nil {
		return nil, err
	}
	defer c.close()

	data, err := c.findJSONByParam("id", id)
	if err != nil {
		return nil, err
	}
	var url models.URL
	err = json.Unmarshal(data, &url)
	if err != nil {
		return nil, err
	}

	return &url, err
}
