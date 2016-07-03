package util

import (
    "testing"
)

func TestParseIP(t *testing.T) {

  val, err := ParseIP("123.456.789.987")

  if val != "" {
    t.Error("Expected \"\" got \"" + val + "\"")
  }

  if err == nil {
    t.Error("Expected error got nil")
  }

  val2, err2 := ParseIP("123.456.789.987:8080")

  if val2 != "123.456.789.987" {
    t.Error("Expected \"123.456.789.987\" got \"" + val2 + "\"")
  }

  if err2 != nil {
    t.Error("Expected error to be nil got error")
  }

}
