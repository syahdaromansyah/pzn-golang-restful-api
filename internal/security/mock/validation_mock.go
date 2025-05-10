package security

import "github.com/stretchr/testify/mock"

type validationMock struct {
	Mock *mock.Mock
}

func NewValidationMock() *validationMock {
	return &validationMock{
		Mock: new(mock.Mock),
	}
}

func (v *validationMock) Struct(s any) error {
	args := v.Mock.Called(s)
	return args.Error(0)
}
