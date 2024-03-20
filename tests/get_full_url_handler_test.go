package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/storage"
)

func noRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func TestGetFullURLHandler(t *testing.T) {
	a := app.Application{}
	err := storage.Init()
	require.NoError(t, err)
	ts := httptest.NewServer(a.Router())
	defer ts.Close()
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
			name: "Positive tests",
			args: args{
				request: "/id",
				urls: map[string]storage.URL{
					"id": {
						ID:      "id",
						FullURL: "http://test.com",
					},
				},
			},
			want: want{
				code:     307,
				location: "http://test.com",
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
			name: "id not presented",
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
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, url := range test.args.urls {
				storage.Save(url)
			}
			req, _ := http.NewRequest(http.MethodGet, ts.URL+test.args.request, nil)
			client := &http.Client{
				CheckRedirect: noRedirect,
			}
			res, err := client.Do(req)
			require.NoError(t, err)
			defer res.Body.Close()
			assert.Equal(t, test.want.code, res.StatusCode)
			assert.Equal(t, test.want.location, res.Header.Get("Location"))
		})
	}
}
