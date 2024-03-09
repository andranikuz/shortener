package tests

import (
	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRequest(t *testing.T, ts *httptest.Server, method,
	path string) *http.Response {
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	return resp
}

func TestGetFullURLHandler(t *testing.T) {
	app := app.Application{}
	app.Init()
	ts := httptest.NewServer(app.Router())
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
					"id1": {
						ID:      "id",
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
	}
	for _, test := range tests {
		for _, url := range test.args.urls {
			storage.Save(url)
		}
		res := testRequest(t, ts, http.MethodGet, test.args.request)
		assert.Equal(t, test.want.code, res.StatusCode)
		assert.Equal(t, test.want.location, res.Header.Get("Location"))
	}
}
