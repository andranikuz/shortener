package shortener

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/storage/memory"
)

func getShortener(t *testing.T) *Shortener {
	storage, err := memory.NewMemoryStorage()
	require.NoError(t, err)
	s := NewShortener(storage)

	return s
}

func TestGenerateShortURL(t *testing.T) {
	s := getShortener(t)
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
	s := getShortener(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, url := range test.args.urls {
				s.storage.Save(context.Background(), url)
			}
			assert.Equal(t, test.want, s.GetFullURL(context.Background(), test.args.id), "GetFullURL(%v)", test.args.id)
		})
	}
}
