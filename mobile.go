package validatorx

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// IsMobile 是否是合法的手机号
func IsMobile(fl validator.FieldLevel) bool {
	return regexp.MustCompile(Mobile).MatchString(fl.Field().String())
}
