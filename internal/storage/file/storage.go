package file

import (
	"context"
	"encoding/json"
	"os"

	"github.com/andranikuz/shortener/internal/models"
)

// FileStorage файловый репозиторий.
type FileStorage struct {
	filePath string
}

// NewFileStorage функция инициализации FileStorage.
func NewFileStorage(path string) (*FileStorage, error) {
	storage := FileStorage{path}
	file, err := os.OpenFile(storage.filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	file.Close()

	return &storage, nil
}

// Save метод сохранения URL.
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

// Get метод получения URL по идентификатору.
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

// GetByFullURL метод получения URL по послной ссылке.
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

// SaveBatch метод сохранения массива URL.
func (storage *FileStorage) SaveBatch(ctx context.Context, urls []models.URL) error {
	for _, url := range urls {
		if err := storage.Save(ctx, url); err != nil {
			return err
		}
	}

	return nil
}

// GetByUserID метод полуения списка URL по userID.
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

// DeleteURLs метод удаления массива URLs.
func (storage *FileStorage) DeleteURLs(ctx context.Context, ids []string, userID string) error {
	return nil
}

// Ping метод проверки статуса соединения.
func (storage FileStorage) Ping() error {
	return nil
}

// GetUsersCount метод получения количества пользователей.
func (storage FileStorage) GetUsersCount(ctx context.Context) (int, error) {
	return 0, nil
}

// GetURLsCount метод получения количества записей.
func (storage FileStorage) GetURLsCount(ctx context.Context) (int, error) {
	return 0, nil
}
