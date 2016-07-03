package qswiss

import (
    "testing"
)

func TestCheckPhoneNumberFormatMinPhoneNoLength(t *testing.T) {
  phoneNumber := "123"
  check := CheckPhoneNumberFormat(phoneNumber)

  if check != false {
    t.Error("Expected phone number 123 to return false")
  }
}

func TestCheckPhoneNumberNoPlusSign(t *testing.T) {
  phoneNumber := "1234556789102"
  check := CheckPhoneNumberFormat(phoneNumber)

  if check != false {
    t.Error("Expected phone number 1234556789102 to return false")
  }
}

func TestCheckPhoneNumberPass(t *testing.T) {
  phoneNumber := "+1234556789102"
  check := CheckPhoneNumberFormat(phoneNumber)

  if check != true {
    t.Error("Expected phone number +1234556789102 to return true")
  }
}
