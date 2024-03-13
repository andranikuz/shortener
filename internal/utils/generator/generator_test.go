package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUniqueID(t *testing.T) {
	assert.NotEqual(t, GenerateUniqueID(), GenerateUniqueID(), GenerateUniqueID(), GenerateUniqueID())
}
