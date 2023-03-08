package echo_validator

import "github.com/go-playground/validator/v10"

type Validate struct {
	validate *validator.Validate
}

func New() *Validate {
	validate := validator.New()
	registerValidation(validate)

	return &Validate{validate: validate}
}

func (v *Validate) Validate(i any) error {
	return v.validate.Struct(i)
}

func registerValidation(validate *validator.Validate) {

}
