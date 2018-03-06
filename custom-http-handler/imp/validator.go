package imp

import "github.com/dimiro1/experiments/custom-http-handler/validator"

type Validator struct{}

func (Validator) Validate(v validator.Validatable) (bool, error) {
	return v.IsValid()
}
