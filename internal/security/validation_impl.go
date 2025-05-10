package security

import "github.com/go-playground/validator/v10"

type validationImpl struct {
	Validator *validator.Validate
}

func NewValidationImpl() Validation {
	return &validationImpl{
		Validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (v *validationImpl) Struct(s any) error {
	return v.Validator.Struct(s)
}
