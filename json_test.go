package util

import (
	"testing"
)

func TestGetJson(t *testing.T) {
	params := map[string]interface{}{"status": "ko", "isecode": "code"}
	jsonByteArray := GetJSON(params)
	s := string(jsonByteArray[:])

	expected := "{\n" +
		"    \"isecode\": \"code\",\n" +
		"    \"status\": \"ko\"\n" +
		"}"

	if s != expected {
		t.Error("Expected wrong json format")
	}

}
