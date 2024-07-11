package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/andranikuz/shortener/internal/config"
)

func TestGenerateShortUrlHandler(t *testing.T) {
	h := getHTTPHandler(t)
	ts := httptest.NewServer(h.Router())
	type want struct {
		code     int
		response string
	}
	type args struct {
		request string
		url     string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Positive test short url",
			args: args{
				request: "/",
				url:     "http://goo/short",
			},
			want: want{
				code:     201,
				response: config.Config.BaseURL,
			},
		},
		{
			name: "Positive test long url",
			args: args{
				request: "/",
				url:     "http://googlegooglegooglegooglegooglegooglegooglegooglegooglegooglegooglegoogle/bigbigbigbigbig",
			},
			want: want{
				code:     201,
				response: config.Config.BaseURL,
			},
		},
		{
			name: "Wrong url",
			args: args{
				request: "/",
				url:     "",
			},
			want: want{
				code:     400,
				response: "",
			},
		},
		{
			name: "Wrong request url",
			args: args{
				request: "/abcd",
				url:     "http://goo/short",
			},
			want: want{
				code:     400,
				response: "",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := strings.NewReader(test.args.url)
			req, _ := http.NewRequest(http.MethodPost, ts.URL+test.args.request, reader)
			res, err := ts.Client().Do(req)
			require.NoError(t, err)
			defer res.Body.Close()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			assert.Contains(t, string(resBody), test.want.response)
		})
	}
}
