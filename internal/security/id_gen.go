package security

type IdGenerator interface {
	Generate(length int) (string, error)
	GenerateCustom(customChars string, length int) (string, error)
}
