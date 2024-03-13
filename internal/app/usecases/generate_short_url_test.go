package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andranikuz/shortener/internal/storage"
)

func TestGenerateShortURL(t *testing.T) {
	storage.Init()
	assert.NotEmpty(t, GenerateShortURL("google.com"))
}
