package security

import (
	"testing"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/model"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/security"

	"github.com/stretchr/testify/assert"
)

func TestStruct(t *testing.T) {
	assert.Nil(t, security.NewValidationImpl().Struct(&model.CreateCategoryRequest{
		Name: "Tools",
	}))
}
