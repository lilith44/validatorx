package validatorx

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

// StartsWithAlpha 以字母开头
func StartsWithAlpha(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	if str == "" {
		return false
	}

	return unicode.IsLetter(rune(str[0]))
}
