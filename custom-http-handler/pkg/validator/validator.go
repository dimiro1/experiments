package validator

type Validatable interface {
	IsValid() (bool, error)
}

type Validator interface {
	Validate(v Validatable) (bool, error)
}
