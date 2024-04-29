package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/andranikuz/shortener/internal/api/rest"
	"github.com/andranikuz/shortener/internal/container"
)

func getHTTPHandler(t *testing.T) rest.HTTPHandler {
	cnt, err := container.NewContainer()
	require.NoError(t, err)
	httpRouter := rest.NewHTTPHandler(cnt)

	return httpRouter
}
