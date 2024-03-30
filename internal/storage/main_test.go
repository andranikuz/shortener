package storage

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andranikuz/shortener/internal/models"
)

func TestSave(t *testing.T) {
	Init()
	url := models.URL{
		ID:      "id",
		FullURL: "url",
	}
	Save(url)
	result, err := Get(url.ID)
	assert.Nil(t, err)
	assert.Equal(t, url.FullURL, result.FullURL)
}

func TestGet(t *testing.T) {
	type args struct {
		id   string
		urls map[string]models.URL
	}
	tests := []struct {
		name    string
		args    args
		want    *models.URL
		wantErr bool
	}{
		{
			name: "positive test",
			args: args{
				id: "id1",
				urls: map[string]models.URL{
					"id1": {
						ID:      "id1",
						FullURL: "http://google.com",
					},
				},
			},
			want: &models.URL{
				ID:      "id1",
				FullURL: "http://google.com",
			},
			wantErr: false,
		},
		{
			name: "wrong id",
			args: args{
				id: "wrong_id",
				urls: map[string]models.URL{
					"id1": {
						ID:      "id1",
						FullURL: "http://google.com",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, url := range tt.args.urls {
				Save(url)
			}
			got, err := Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
