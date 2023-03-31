package validatorx

import "github.com/go-playground/validator/v10"

type EchoValidator struct {
	validate *validator.Validate
}

func NewEchoValidator(validate *validator.Validate) *EchoValidator {
	return &EchoValidator{validate: validate}
}

func (ev *EchoValidator) Validate(i any) error {
	return ev.validate.Struct(i)
}
