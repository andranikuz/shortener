package tests

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/andranikuz/shortener/internal/api/rest"
	"github.com/andranikuz/shortener/internal/container"
)

func ExampleGenerateShortURLHandler() {
	cnt, _ := container.NewContainer()
	h := rest.NewHTTPHandler(cnt)
	ts := httptest.NewServer(h.Router(context.Background()))
	reader := strings.NewReader("google.com")
	req, _ := http.NewRequest(http.MethodPost, ts.URL, reader)
	res, _ := ts.Client().Do(req)
	defer res.Body.Close()
	fmt.Println("Status Code:", res.StatusCode)

	// Output:
	// Status Code: 201
}

func ExampleGenerateShortUrlJSONHandler() {
	cnt, _ := container.NewContainer()
	h := rest.NewHTTPHandler(cnt)
	ts := httptest.NewServer(h.Router(context.Background()))
	reader := strings.NewReader("{\"url\": \"http://google.com\"}")
	req, _ := http.NewRequest(http.MethodPost, ts.URL, reader)
	res, _ := ts.Client().Do(req)
	defer res.Body.Close()
	fmt.Println("Status Code:", res.StatusCode)

	// Output:
	// Status Code: 201
}

func ExampleDeleteURLsHandler() {
	cnt, _ := container.NewContainer()
	h := rest.NewHTTPHandler(cnt)
	ts := httptest.NewServer(h.Router(context.Background()))
	reader := strings.NewReader("[\"id\"]")
	req, _ := http.NewRequest(http.MethodDelete, ts.URL+"/api/user/urls", reader)
	res, _ := ts.Client().Do(req)
	defer res.Body.Close()
	fmt.Println("Status Code:", res.StatusCode)
	
	// Output:
	// Status Code: 202
}
