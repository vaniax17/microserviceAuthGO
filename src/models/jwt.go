package models

// This package defines the data models used throughout the microservice.

import "github.com/golang-jwt/jwt/v5"

// Claims represents the custom claims structure for JWT. It embeds jwt.RegisteredClaims
// for standard JWT claims and adds a Username field for the subject.
type Claims struct {
	Username string `json:"sub"` // Username of the user (subject of the JWT)
	jwt.RegisteredClaims       // Standard JWT claims (e.g., expiry, audience)
}
