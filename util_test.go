package qswiss

import (
	"testing"
)

func TestIsNil(t *testing.T) {

	if IsNil("") != true {
		t.Error("Expected empty string to return true from IsNil")
	}

	if IsNil("string") != false {
		t.Error("Expected \"string\" to return false from IsNil")
	}

}

func TestCheckUsername(t *testing.T) {
	if CheckUsername("v", "6", "10") != false {
		t.Error("Expected v username, min 6, max 10 to return false (too short)")
	}
	if CheckUsername("usernameusernameusername", "6", "10") != false {
		t.Error("Expected usernameusernameusername username, min 6, max 10 to return false (too long)")
	}
	if CheckUsername("someu*$@", "6", "10") != false {
		t.Error("Expected someu*$@ username, min 6, max 10 to return false (special characters)")
	}
	if CheckUsername("johndoe", "6", "10") != true {
		t.Error("Expected johndoe username, min 6, max 10 to return true (should pass)")
	}
	if CheckUsername("johnDoe", "6", "10") != true {
		t.Error("Expected johnDoe username, min 6, max 10 to return true (should pass)")
	}
}

func TestIsTooShort(t *testing.T) {
	if IsTooShort("joe", 4) != true {
		t.Error("Expected joe string, length 4 to return true (should pass because it is too short)")
	}
	if IsTooShort("johndoe", 4) != false {
		t.Error("Expected johndoe string, length 4 to return true (should fail because it is not too short)")
	}
}

func TestEmptyStringConvertor(t *testing.T) {
	if EmptyStringConvertor("") != nil {
		t.Error("Expected empty string to be returned as nil")
	}

}

func TestStringInSlice(t *testing.T) {
	if StringInSlice("Apple", []string{"Banana", "Orange", "Blackberry"}) != false {
		t.Error("Expected false (slice doesn't contain string)")
	}

	if StringInSlice("Apple", []string{"Banana", "Orange", "Apple", "Blackberry"}) != true {
		t.Error("Expected true (slice contain string)")
	}
}

func TestSetDefaultIfNil(t *testing.T) {
	val := SetDefaultIfNil("", "Go Home")
	if val != "Go Home" {
		t.Error("Expected to get \"Go Home\" string got \"" + val + "\"")
	}

	val2 := SetDefaultIfNil("Go to cinema", "Go Home")
	if val2 != "Go to cinema" {
		t.Error("Expected to get \"Go to cinema\" string got \"" + val2 + "\"")
	}
}

func TestUrlEncoded(t *testing.T) {

	val, err := URLEncoded("$$$£@$!%^&^T&^@:≈≈≈≈")

	if val != "" {
		t.Error("Expected to get empty string, got " + val)
	}

	if err == nil {
		t.Error("Expected to get error")
	}

	val2, err2 := URLEncoded("index_page.html")

	if val2 != "index_page.html" {
		t.Error("Expected index_page.html got " + val2)
	}

	if err2 != nil {
		t.Error("Expected error to be nil")
	}

}
