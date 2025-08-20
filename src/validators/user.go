package validators

// This package provides functions for validating user-related data, such as email addresses and usernames.
// It utilizes the `govalidator` library for common validation tasks.

import "github.com/asaskevich/govalidator"

// IsEmail validates if the provided string is a valid email address.
// It uses the govalidator.IsEmail function for the validation.
// Returns true if the email is valid, false otherwise.
func IsEmail(email string) bool {
	return govalidator.IsEmail(email)
}

// IsCorrectUsername validates if the provided username meets the length requirements.
// A valid username must have a length between 3 and 30 characters (inclusive).
// Returns true if the username is within the correct length, false otherwise.
func IsCorrectUsername(username string) bool {
	if len(username) < 3 && len(username) > 30 {
		return false
	}

	return true
}
