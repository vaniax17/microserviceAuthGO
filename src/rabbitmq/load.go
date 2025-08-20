package rabbitmq

// This package handles the initialization of the RabbitMQ publisher for the application.
// It sets up the connection to RabbitMQ and provides a global publisher instance.

import (
	"log"
	"os"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
)

var (
	Pub *amqp.Publisher // Pub is the global AMQP publisher instance used for sending messages.
)

// Init initializes the RabbitMQ publisher.
// It retrieves the AMQP URI from environment variables, configures the Watermill logger,
// sets up the AMQP pub/sub configuration for durable queues, and creates the publisher instance.
// If any step fails, it logs a fatal error and exits.
func Init() {
	// Retrieve the AMQP URI from environment variables.
	amqpURI := os.Getenv("AMQP_URI")
	// Create a new standard logger for Watermill, enabling debug and tracing logs.
	logger := watermill.NewStdLogger(true, true)
	// Configure the AMQP pub/sub with a durable setup and a custom queue naming function.
	amqpConfig := amqp.NewDurablePubSubConfig(
		amqpURI, func(topic string) string {
			return "queue_" + topic // Defines the queue name based on the topic.
		})

	// Create a new AMQP publisher instance.
	publisher, err := amqp.NewPublisher(amqpConfig, logger)
	if err != nil {
		log.Fatalf("Failed to create Publisher: %s", err)
	}

	Pub = publisher

}
