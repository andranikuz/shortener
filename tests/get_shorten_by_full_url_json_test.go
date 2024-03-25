package tests

import (
	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetShortenByFullUrlJSONHandler(t *testing.T) {
	a := app.Application{}
	err := storage.Init()
	require.NoError(t, err)
	ts := httptest.NewServer(a.Router())
	type args struct {
		body    string
		urls    map[string]models.URL
		request string
	}
	type want struct {
		code int
		host string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "positive test",
			args: args{
				body: "{\"url\": \"http://google.com\"}",
				urls: map[string]models.URL{
					"id": {
						ID:      "id",
						FullURL: "http://google.com",
					},
				},
				request: "/api/shorten",
			},
			want: want{
				code: 201,
				host: config.Config.BaseURL,
			},
		},
		{
			name: "wrong request",
			args: args{
				body: "{\"url\": \"http://google.com\"}",
				urls: map[string]models.URL{
					"id": {
						ID:      "id",
						FullURL: "http://google.com",
					},
				},
				request: "/api/wrong",
			},
			want: want{
				code: 404,
			},
		},
		{
			name: "wrong json",
			args: args{
				body: "\"url\": \"http://google.com\"}",
				urls: map[string]models.URL{
					"id": {
						ID:      "id",
						FullURL: "http://google.com",
					},
				},
				request: "/api/wrong",
			},
			want: want{
				code: 404,
			},
		},
		{

			name: "not found",
			args: args{
				body:    "{\"url\": \"http://yandex.com\"}",
				urls:    map[string]models.URL{},
				request: "/api/wrong",
			},
			want: want{
				code: 404,
				host: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, url := range tt.args.urls {
				storage.Save(url)
			}
			reader := strings.NewReader(tt.args.body)
			req, _ := http.NewRequest(http.MethodPost, ts.URL+tt.args.request, reader)
			res, err := ts.Client().Do(req)
			defer res.Body.Close()

			require.NoError(t, err)
			assert.Equal(t, tt.want.code, res.StatusCode)
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			assert.Contains(t, string(resBody), tt.want.host)
		})
	}
}
