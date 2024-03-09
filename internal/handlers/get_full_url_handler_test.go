package handlers

import (
	"github.com/andranikuz/shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFullURLHandler(t *testing.T) {
	type want struct {
		code     int
		location string
	}
	type args struct {
		request string
		urls    map[string]storage.URL
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Positive test",
			args: args{
				request: "/id1",
				urls: map[string]storage.URL{
					"id1": {
						ID:      "id1",
						FullURL: "http://google.com",
					},
				},
			},
			want: want{
				code:     307,
				location: "http://google.com",
			},
		},
		{
			name: "id not found",
			args: args{
				request: "/id2",
				urls:    map[string]storage.URL{},
			},
			want: want{
				code:     400,
				location: "",
			},
		},
		{
			name: "bad request",
			args: args{
				request: "/",
				urls:    map[string]storage.URL{},
			},
			want: want{
				code:     400,
				location: "",
			},
		},
	}
	storage.Init()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, url := range test.args.urls {
				storage.Save(url)
			}
			request := httptest.NewRequest(http.MethodGet, test.args.request, nil)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			GetFullURLHandler(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			assert.Equal(t, test.want.location, res.Header.Get("Location"))
		})
	}
}
