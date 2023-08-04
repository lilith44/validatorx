package validatorx

import "github.com/go-playground/validator/v10"

func New() (*validator.Validate, error) {
	validate := validator.New()
	if err := registerValidation(validate); err != nil {
		return nil, err
	}
	return validate, nil
}

func registerValidation(validate *validator.Validate) error {
	register := map[string]validator.Func{
		"chinese":    ContainsOnlyChinese,
		"id_card":    IsIdCard,
		"id_card_15": IsIdCardWith15Len,
		"mobile":     IsMobile,
		"password":   IsStrongPassword,
		"password_with_lower_upper_digit_special": IsVeryStrongPassword,
		"starts_with_alpha":                       StartsWithAlpha,
	}

	for tag, fun := range register {
		if err := validate.RegisterValidation(tag, fun); err != nil {
			return err
		}
	}
	return nil
}
