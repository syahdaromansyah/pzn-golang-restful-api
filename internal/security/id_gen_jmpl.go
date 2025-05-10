package security

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type idGenImpl struct{}

func NewIdGenImpl() IdGenerator {
	return new(idGenImpl)
}

func (i *idGenImpl) Generate(length int) (string, error) {
	id, err := gonanoid.New(length)
	return id, err
}

func (i *idGenImpl) GenerateCustom(customChars string, length int) (string, error) {
	id, err := gonanoid.Generate(customChars, length)
	return id, err
}
