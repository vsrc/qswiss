package qswiss

import (
	"net/url"
	"regexp"
	"strings"
)

// DateTimeFormat defines accepted format for datetime
var DateTimeFormat = "2006-1-2 15:04:05"

// IsNil checks if string is empty or contains only whitespace
func IsNil(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return true
	}

	return false
}

// GetReservedKeywords returns array of system reserved keywords
// useful in projects where users are allowed naming (like setting username)
func GetReservedKeywords() []string {
	return []string{"user", "post", "username", "password", "status",
		"table", "admin", "administrator"}
}

// CheckUsername check if provided string complies with set rules
// accepts username string, minimum number of characters and maximal number of characters
// returns boolean, true if string complies with rules, false otherwise
func CheckUsername(username string, min string, max string) bool {
	matched, err := regexp.MatchString("^[a-zA-Z.0-9_-]{"+min+","+max+"}$", username)
	PanicIf(err)
	return matched
}

// IsTooShort checks if provided string is too short
// if number of characters in string is lower than in provides length
func IsTooShort(s string, length int) bool {
	if len(s) < length {
		return true
	}

	return false
}

// EmptyStringConvertor converts empty string to nil
// if provided string is not empty converts it to array of byte
func EmptyStringConvertor(s string) []byte {
	if !IsNil(s) {
		return []byte(s)
	}

	return nil
}

// StringInSlice checks if provides array of string contains provided string
// accepts string for search and array of strings to check
// returns boolean, true if there is a match, false otherwise
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// SetDefaultIfNil returns default value if provided string is empty
// accepts value string to check and default string
// returns result string
func SetDefaultIfNil(v string, d string) string {
	if IsNil(v) {
		return d
	}

	return v
}

// URLEncoded converts string into url friendly string
// accepts string to be converted
// returns converted string and nil if conversion happend
// returns empty string and error if error happend
func URLEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
