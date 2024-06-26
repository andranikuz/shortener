package tests

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/andranikuz/shortener/internal/config"
)

func TestGetShortenByFullUrlJSONHandler(t *testing.T) {
	h := getHTTPHandler(t)
	ts := httptest.NewServer(h.Router(context.Background()))
	type args struct {
		body    string
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
				body:    "{\"url\": \"http://google.com\"}",
				request: "/api/shorten",
			},
			want: want{
				code: 201,
				host: config.Config.BaseURL,
			},
		},
		{
			name: "wrong json",
			args: args{
				body:    "\"url\": \"http://google.com\"}",
				request: "/api/shorten",
			},
			want: want{
				code: 400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.args.body)
			req, _ := http.NewRequest(http.MethodPost, ts.URL+tt.args.request, reader)
			res, err := ts.Client().Do(req)

			require.NoError(t, err)
			defer res.Body.Close()
			assert.Equal(t, tt.want.code, res.StatusCode)
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			assert.Contains(t, string(resBody), tt.want.host)
		})
	}
}
