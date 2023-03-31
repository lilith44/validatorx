package validatorx

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ContainsOnlyChinese 是否仅包含中文字符
func ContainsOnlyChinese(fl validator.FieldLevel) bool {
	return regexp.MustCompile(OnlyChinese).MatchString(fl.Field().String())
}
