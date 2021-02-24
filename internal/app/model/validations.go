package model

import validation "github.com/go-ozzo/ozzo-validation"

func requiredIf(b bool) validation.RuleFunc {
	return func(value interface{}) error {
		if b {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}
