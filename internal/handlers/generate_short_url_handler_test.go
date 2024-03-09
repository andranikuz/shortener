package handlers

import (
	"github.com/andranikuz/shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGenerateShortUrlHandler(t *testing.T) {
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
				response: "http://localhost:8080/",
			},
		},
		{
			name: "Positive test short url",
			args: args{
				request: "/",
				url:     "http://googlegooglegooglegooglegooglegooglegooglegooglegooglegooglegooglegoogle/bigbigbigbigbig",
			},
			want: want{
				code:     201,
				response: "http://localhost:8080/",
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
	storage.Init()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := strings.NewReader(test.args.url)
			request := httptest.NewRequest(http.MethodPost, test.args.request, reader)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			GenerateShortURLHandler(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.Contains(t, string(resBody), test.want.response)
		})
	}
}
