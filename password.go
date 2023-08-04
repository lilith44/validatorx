package validatorx

import "github.com/go-playground/validator/v10"

// IsStrongPassword 是否是强密码
// 指同时包含数字，字母，特殊字符的字符串
func IsStrongPassword(fl validator.FieldLevel) bool {
	var digit, letter, special bool
	for _, r := range []rune(fl.Field().String()) {
		if r < 32 || r > 126 && r < 256 {
			return false
		}

		if r >= '0' && r <= '9' {
			digit = true
		} else if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			letter = true
		} else {
			special = true
		}

		if digit && letter && special {
			return true
		}
	}
	return false
}

// IsVeryStrongPassword 是否是强密码
// 指同时包含数字，大写字母，小写字母以及特殊字符的字符串
func IsVeryStrongPassword(fl validator.FieldLevel) bool {
	var digit, lower, upper, special bool
	for _, r := range []rune(fl.Field().String()) {
		if r < 32 || r > 126 && r < 256 {
			return false
		}

		if r >= '0' && r <= '9' {
			digit = true
		} else if r >= 'a' && r <= 'z' {
			lower = true
		} else if r >= 'A' && r <= 'Z' {
			upper = true
		} else {
			special = true
		}

		if digit && lower && upper && special {
			return true
		}
	}
	return false
}
