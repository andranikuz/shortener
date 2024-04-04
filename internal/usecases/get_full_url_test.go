package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/models"
)

func TestGetFullURL(t *testing.T) {
	a, err := app.NewApplication()
	require.NoError(t, err)
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
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, url := range test.args.urls {
				a.DB.Save(url)
			}
			assert.Equal(t, test.want, GetFullURL(*a, test.args.id), "GetFullURL(%v)", test.args.id)
		})
	}
}
