package core

// This package handles the creation and parsing of JSON Web Tokens (JWTs).
// It provides functions to generate a new JWT for a given username and to parse
// and validate an existing JWT, extracting its claims.

import (
	"log"
	"microserviceAuthGO/src/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// secretJwtKey is the secret key used for signing and verifying JWTs.
// It is loaded from the environment variable "SECRET_JWT_KEY".
var secretJwtKey = []byte(os.Getenv("SECRET_JWT_KEY"))

// CreateJWT generates a new JWT for the given username.
// The token includes the username as a subject, issue time (iat), and expiration time (exp).
// It is signed using HS256 algorithm with the secretJwtKey.
// Returns the signed JWT string and an error if creation fails.
func CreateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(secretJwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseClaims parses and validates a JWT string and extracts its claims.
// It expects the token to be signed with the secretJwtKey and to contain claims
// that conform to the models.Claims structure.
// Returns a pointer to the Claims struct if successful, otherwise logs a fatal error and exits.
func ParseClaims(tokenString string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (any, error) {
		return secretJwtKey, nil
	})
	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*models.Claims); ok {
		return claims, nil
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
	return nil, nil // Should not be reached due to log.Fatal
}
