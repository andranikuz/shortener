package usecases

import (
	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetURLByFullURL(t *testing.T) {
	type args struct {
		fullURL string
		urls    map[string]models.URL
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "positive test",
			args: args{
				fullURL: "http://google.com",
				urls: map[string]models.URL{
					"id": {
						ID:      "id",
						FullURL: "http://google.com",
					},
				},
			},
			want:    "http://google.com",
			wantErr: false,
		},
		{
			name: "negative test",
			args: args{
				fullURL: "http://yandex.ru",
			},
			want:    "",
			wantErr: true,
		},
	}
	storage.Init()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, url := range tt.args.urls {
				storage.Save(url)
			}
			got, err := GetURLByFullURL(tt.args.fullURL)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, got.FullURL)
				assert.NoError(t, err)
			}
		})
	}
}
