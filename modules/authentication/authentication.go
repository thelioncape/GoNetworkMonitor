package authentication

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// CredentialsUnhashed is a type that represents a user's credentials (plaintext password)
type CredentialsUnhashed struct {
	Username string
	Password string
}

// CredentialsHashed is a type that represents a user's credentials (hashed password)
type CredentialsHashed struct {
	Username   string
	Hashedpass string
}

// HashPass hashes the given password using the cost defined in this function and returns the hash as []byte
func HashPass(password []byte) []byte {
	cost := 10
	hash, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}

func CheckCreds(creds CredentialsUnhashed) bool {
	return true
}
