package util

import "regexp"

func ValidatePassword(password string) bool {
	patterns := []string{".{8,}", "[a-z]", "[A-Z]", "[0-9]"}
	for _, pattern := range patterns {
		match, _ := regexp.MatchString(pattern, password)
		if !match {
			return false
		}
	}
	return true
}
