package helpers

import "unicode"

func PasswordIsValid(password string) bool {
	hasCorrectLength := false
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	if len(password) >= 8 && len(password) <= 24 {
		hasCorrectLength = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasCorrectLength && hasUpper && hasLower && hasNumber && hasSpecial
}
