package core

// This package provides cryptographic functionalities for hashing and comparing passwords.
// It uses the argon2id algorithm for secure password storage.

import (
	"log"

	"github.com/alexedwards/argon2id"
)

// Hash generates a secure hash of the provided password using argon2id.
// It returns the hashed password as a string. If an error occurs during hashing,
// it logs a fatal error and exits.
func Hash(password string) string {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatalf("Error creating hash: %v", err)
	}

	return hash
}

// Compare compares a plaintext password with a hashed password sum.
// It returns true if the password matches the hash, and false otherwise.
// If an error occurs during comparison, it logs a fatal error and exits.
func Compare(password, passwordHashSum string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, passwordHashSum)
	if err != nil {
		log.Fatalf("Error comparing password: %v", err)
	}
	return match
}
