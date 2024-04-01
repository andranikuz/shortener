package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/models"
)

func TestGetFullURL(t *testing.T) {
	a := app.Application{}
	err := a.Init()
	require.NoError(t, err)
	app.App = &a
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
				app.App.DB.Save(url)
			}
			assert.Equal(t, test.want, GetFullURL(test.args.id), "GetFullURL(%v)", test.args.id)
		})
	}
}
