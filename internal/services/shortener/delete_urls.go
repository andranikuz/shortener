package shortener

import (
	"context"
)

func (s *Shortener) DeleteURLs(ctx context.Context, ids []string, userID string) {
	go s.storage.DeleteURLs(ctx, ids, userID)
}
