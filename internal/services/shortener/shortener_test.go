package shortener

import (
	"context"
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/stretchr/testify/assert"

	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/storage/memory"
)

func getShortener() *Shortener {
	storage, _ := memory.NewMemoryStorage()
	s := NewShortener(storage)

	return s
}

func TestGenerateShortURL(t *testing.T) {
	s := getShortener()
	shorter, err := s.GenerateShortURL(context.Background(), "google.com", "userID")
	assert.NoError(t, err)
	assert.NotEmpty(t, shorter)
}

func TestGetFullURL(t *testing.T) {
	type args struct {
		urls map[string]models.URL
		id   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "positive test",
			args: args{
				urls: map[string]models.URL{
					"id1": {
						ID:      "id1",
						FullURL: "http://google.com",
						UserID:  "userId",
					},
				},
				id: "id1",
			},
			want: "http://google.com",
		},
		{
			name: "negative test",
			args: args{
				urls: map[string]models.URL{},
				id:   "id2",
			},
			want: "",
		},
	}
	s := getShortener()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, url := range test.args.urls {
				s.storage.Save(context.Background(), url)
			}
			fullURL, _ := s.GetFullURL(context.Background(), test.args.id)
			assert.Equal(t, test.want, fullURL, "GetFullURL(%v)", test.args.id)
		})
	}
}

func BenchmarkGenerateShortURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := getShortener()
		s.GenerateShortURL(context.Background(), "google.com", "userID")
	}
}

func BenchmarkGetFullURL(b *testing.B) {
	s := getShortener()
	s.storage.Save(context.Background(), models.URL{
		ID:      "id",
		FullURL: "http://google.com",
		UserID:  "userId",
	})
	for i := 0; i < b.N; i++ {
		s.GetFullURL(context.Background(), "id")
	}
}

func BenchmarkDeleteURLs(b *testing.B) {
	s := getShortener()
	userID, _ := uuid.GenerateUUID()
	s.storage.Save(context.Background(), models.URL{
		ID:      "id",
		FullURL: "http://google.com",
		UserID:  userID,
	})
	var ids []string
	ids = append(ids, "id")
	for i := 0; i < b.N; i++ {
		s.DeleteURLs(context.Background(), ids, userID)
	}
}

func BenchmarkGetUserURLs(b *testing.B) {
	s := getShortener()
	userID := "userId"
	s.storage.Save(context.Background(), models.URL{
		ID:      "id",
		FullURL: "http://google.com",
		UserID:  userID,
	})
	for i := 0; i < b.N; i++ {
		s.GetUserURLs(context.Background(), userID)
	}
}

func BenchmarkGenerateShortURLBatch(b *testing.B) {
	s := getShortener()
	userID := "userId"
	var items []OriginalItem
	items = append(items,
		OriginalItem{
			CorrelationID: "id",
			OriginalURL:   "url",
		},
		OriginalItem{
			CorrelationID: "id1",
			OriginalURL:   "url2",
		},
	)
	for i := 0; i < b.N; i++ {
		s.GenerateShortURLBatch(context.Background(), items, userID)
	}
}
