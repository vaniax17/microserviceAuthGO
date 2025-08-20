package models

// This package defines the data models for the application's database entities.

// User represents a user in the system. It is mapped to a database table by GORM.
type User struct {
	Id             int64  `gorm:"primary_key"`               // Unique identifier for the user, serves as the primary key.
	Username       string `gorm:"unique;not null;index:idx_username"` // Unique username, cannot be null, indexed for faster lookups.
	Email          string `gorm:"unique;not null;index:idx_email"`    // Unique email address, cannot be null, indexed for faster lookups.
	HashedPassword string `gorm:"not null;index:idx_hashed_password"` // Hashed password of the user, cannot be null, indexed.
}
