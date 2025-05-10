package security

type Validation interface {
	Struct(s any) error
}
