package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/andranikuz/shortener/internal/app"
)

func TestGenerateShortURL(t *testing.T) {
	a, err := app.NewApplication()
	require.NoError(t, err)
	shorter, err := GenerateShortURL(*a, "google.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, shorter)
}
