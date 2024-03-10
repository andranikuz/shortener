package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateUniqueID(t *testing.T) {
	assert.NotEqual(t, GenerateUniqueID(), GenerateUniqueID(), GenerateUniqueID(), GenerateUniqueID())
}
