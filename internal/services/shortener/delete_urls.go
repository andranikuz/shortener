package shortener

import (
	"context"
)

// DeleteURLs метод удаления массива URLs.
func (s Shortener) DeleteURLs(ids []string, userID string) {
	go s.storage.DeleteURLs(context.Background(), ids, userID)
}
