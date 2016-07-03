package util

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	mathrand "math/rand"
	"time"
)

// GenerateToken returns random hash that can be used as token
func GenerateToken(len int) string {

	base := make([]byte, len)
	_, err := rand.Read(base)
	PanicIf(err)
	salt := base64.StdEncoding.EncodeToString(base)

	hash := sha512.New()
	hash.Write([]byte(salt))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// GenerateRandomInt returns random integer
// where max is maximal value this function can return
func GenerateRandomInt(max int) int {
	mathrand.Seed(time.Now().UnixNano())
	return mathrand.Intn(max)
}
