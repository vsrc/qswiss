package qswiss

import (
    "testing"
    "reflect"
)

func TestGenerateToken(t *testing.T) {
  token := GenerateToken(5)

  if reflect.TypeOf(token).Kind() != reflect.String {
    t.Error("Returned token is not of type string")
  }
}
