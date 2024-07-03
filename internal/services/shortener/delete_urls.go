package shortener

import (
	"context"
)

// DeleteURLs метод удаления массива URLs.
func (s Shortener) DeleteURLs(ctx context.Context, ids []string, userID string) {
	go s.storage.DeleteURLs(ctx, ids, userID)
}
