package security

import "github.com/stretchr/testify/mock"

type idGenMock struct {
	Mock *mock.Mock
}

func NewIdGenMock() *idGenMock {
	return &idGenMock{
		Mock: new(mock.Mock),
	}
}

func (i *idGenMock) Generate(length int) (string, error) {
	args := i.Mock.Called(length)
	return args.String(0), args.Error(1)
}

func (i *idGenMock) GenerateCustom(customChars string, length int) (string, error) {
	args := i.Mock.Called(customChars, length)
	return args.String(0), args.Error(1)
}
