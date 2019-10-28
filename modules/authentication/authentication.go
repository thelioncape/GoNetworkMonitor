package authentication

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPass hashes the given password using the cost defined in this function and returns the hash as []byte
func HashPass(password []byte) []byte {
	cost := 10
	hash, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}
