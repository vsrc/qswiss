package qswiss

import (
	"strings"
)

// allowed length of string in order to be accepted as phone number
const (
	LenOfPhoneNumber    int = 25
	MinLenOfPhoneNumber int = 11
)

// CheckPhoneNumberFormat checks if provided string complies with set of rules
func CheckPhoneNumberFormat(phoneNumber string) bool {
	// ++4412123456 (12 chars), +4412123456 (11 chars), 004412123456 (12 chars),

	if len(phoneNumber) < MinLenOfPhoneNumber {
		return false
	}

	if strings.Contains(phoneNumber, "++") || strings.Contains(phoneNumber, "+") {
		return true
	}

	return false
}
