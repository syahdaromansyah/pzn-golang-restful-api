package security

import (
	"testing"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/security"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSuccess(t *testing.T) {
	id, err := security.NewIdGenImpl().Generate(36)
	assert.Nil(t, err)
	assert.Equal(t, 36, len(id))
}

func TestGenerateCustomSuccess(t *testing.T) {
	id, err := security.NewIdGenImpl().GenerateCustom("abc", 36)
	assert.Nil(t, err)
	assert.Equal(t, 36, len(id))
}
