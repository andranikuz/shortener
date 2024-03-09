package usecases

import (
	"github.com/andranikuz/shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	storage.Init()
	assert.NotEmpty(t, GenerateShortURL("google.com"))
}
