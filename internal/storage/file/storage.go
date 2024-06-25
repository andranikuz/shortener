package file

import (
	"context"
	"encoding/json"
	"os"

	"github.com/andranikuz/shortener/internal/models"
)

type FileStorage struct {
	filePath string
}

func NewFileStorage(path string) (*FileStorage, error) {
	storage := FileStorage{path}
	file, err := os.OpenFile(storage.filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	file.Close()

	return &storage, nil
}

// Save url
func (storage *FileStorage) Save(ctx context.Context, url models.URL) error {
	data, err := json.Marshal(&url)
	if err != nil {
		return err
	}
	p, err := newProducer(storage.filePath)
	if err != nil {
		return err
	}
	if err := p.write(data); err != nil {
		return err
	}

	return nil
}

// Get url
func (storage *FileStorage) Get(ctx context.Context, id string) (*models.URL, error) {
	c, err := newConsumer(storage.filePath)
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

// Get url by full_url
func (storage *FileStorage) GetByFullURL(ctx context.Context, fullURL string) (*models.URL, error) {
	c, err := newConsumer(storage.filePath)
	if err != nil {
		return nil, err
	}
	defer c.close()

	data, err := c.findJSONByParam("full-url", fullURL)
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

// Save batch of urls
func (storage *FileStorage) SaveBatch(ctx context.Context, urls []models.URL) error {
	for _, url := range urls {
		if err := storage.Save(ctx, url); err != nil {
			return err
		}
	}

	return nil
}

func (storage *FileStorage) GetByUserID(ctx context.Context, userID string) ([]models.URL, error) {
	c, err := newConsumer(storage.filePath)
	if err != nil {
		return nil, err
	}
	defer c.close()

	data, err := c.findJSONByParam("user-id", userID)
	if err != nil {
		return nil, err
	}
	var url models.URL
	err = json.Unmarshal(data, &url)
	if err != nil {
		return nil, err
	}
	var urls []models.URL
	urls = append(urls, url)

	return urls, err
}

func (storage *FileStorage) DeleteURLs(ctx context.Context, ids []string, userID string) error {
	return nil
}
