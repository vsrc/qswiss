package qswiss
 
import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	mathrand "math/rand"
	"time"
)

// GenerateToken returns random hash that can be used as token
// accepts len integer as number of characters for autogenterated salt
func GenerateToken(len int) (string, error) {

	base := make([]byte, len)
	_, err := rand.Read(base)
	if err != nil {
		return "", err
	}
	salt := base64.StdEncoding.EncodeToString(base)

	hash := sha512.New()
	hash.Write([]byte(salt))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil)), nil
}

// GenerateRandomInt returns random integer
// where max is maximal value this function can return
func GenerateRandomInt(max int) int {
	mathrand.Seed(time.Now().UnixNano())
	return mathrand.Intn(max)
}
