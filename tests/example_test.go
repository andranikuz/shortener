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

func ExampleHTTPHandler_GenerateShortURLHandler() {
	cnt, _ := container.NewContainer()
	h := rest.NewHTTPHandler(cnt)
	ts := httptest.NewServer(h.Router(context.Background()))
	reader := strings.NewReader("google.com")
	req, err := http.NewRequest(http.MethodPost, ts.URL, reader)
	if err != nil {
		fmt.Println("Status Code:", 500)
	}
	res, err := ts.Client().Do(req)
	if err != nil {
		fmt.Println("Status Code:", 500)
	}
	defer res.Body.Close()
	fmt.Println("Status Code:", res.StatusCode)

	// Output:
	// Status Code: 201
}

func ExampleHTTPHandler_GenerateShortUrlJSONHandler() {
	cnt, _ := container.NewContainer()
	h := rest.NewHTTPHandler(cnt)
	ts := httptest.NewServer(h.Router(context.Background()))
	reader := strings.NewReader("{\"url\": \"http://google.com\"}")
	req, err := http.NewRequest(http.MethodPost, ts.URL, reader)
	if err != nil {
		fmt.Println("Status Code:", 500)
	}
	res, err := ts.Client().Do(req)
	if err != nil {
		fmt.Println("Status Code:", 500)
	}
	defer res.Body.Close()
	fmt.Println("Status Code:", res.StatusCode)

	// Output:
	// Status Code: 201
}

func ExampleHTTPHandler_DeleteURLsHandler() {
	cnt, _ := container.NewContainer()
	h := rest.NewHTTPHandler(cnt)
	ts := httptest.NewServer(h.Router(context.Background()))
	reader := strings.NewReader("[\"id\"]")
	req, err := http.NewRequest(http.MethodDelete, ts.URL+"/api/user/urls", reader)
	if err != nil {
		fmt.Println("Status Code:", 500)
	}
	res, err := ts.Client().Do(req)
	if err != nil {
		fmt.Println("Status Code:", 500)
	}
	defer res.Body.Close()
	fmt.Println("Status Code:", res.StatusCode)

	// Output:
	// Status Code: 202
}
