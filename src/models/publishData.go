package models

// This package defines the data models for messages published to RabbitMQ.

// PublishData represents the structure of data to be published to a message queue.
// It contains user-related information such as Username and Email.
type PublishData struct {
	Username string `json:"username"` // Username of the user
	Email    string `json:"email"`    // Email of the user
}
