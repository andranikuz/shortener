package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/andranikuz/shortener/internal/app"
)

func TestGenerateShortURL(t *testing.T) {
	a := app.Application{}
	err := a.Init()
	require.NoError(t, err)
	assert.NotEmpty(t, GenerateShortURL(a, "google.com"))
}
