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
	app.App = &a
	assert.NotEmpty(t, GenerateShortURL("google.com"))
}
