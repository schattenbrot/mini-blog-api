package utils

import "unicode"

const PASSWORD_MIN_LENGTH = 8
const PASSWORD_MAX_LENGTH = 24

// PasswordIsValid checks for a password validity.
//
// Returns false if:
//   - the password is too short (less than 8 characters)
//   - the password is too long (more than 24 characters)
//   - the password contains no upper case character
//   - the password contains no lower case character
//   - the password contains no number
//   - the password contains no special character
func PasswordIsValid(password string) bool {
	hasCorrectLength := false
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	if passwordLengthIsValid(password) {
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

func passwordLengthIsValid(password string) bool {
	if len(password) < PASSWORD_MIN_LENGTH {
		return false
	}
	if len(password) > PASSWORD_MAX_LENGTH {
		return false
	}
	return true
}
