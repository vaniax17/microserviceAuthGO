package rabbitmq

// This package provides functions for publishing messages to RabbitMQ topics.
// It utilizes the initialized Watermill publisher instance (Pub) from the load package.

import (
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
)

// PublishUserCreate publishes a user creation event message to the "user.create" topic.
// The `data` parameter is a Watermill message containing the user data to be published.
// If the publishing fails, it logs a fatal error.
func PublishUserCreate(data *message.Message) {
	err := Pub.Publish("user.create", data)
	if err != nil {
		log.Fatalf(
			"Failed to publish user create event: %s",
			err.Error())
	}
}
