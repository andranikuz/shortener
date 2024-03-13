package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andranikuz/shortener/internal/storage"
)

func TestGetFullURL(t *testing.T) {
	type args struct {
		urls map[string]storage.URL
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
				urls: map[string]storage.URL{
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
				urls: map[string]storage.URL{},
				id:   "id2",
			},
			want: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, url := range test.args.urls {
				storage.Save(url)
			}
			assert.Equal(t, test.want, GetFullURL(test.args.id), "GetFullURL(%v)", test.args.id)
		})
	}
}
