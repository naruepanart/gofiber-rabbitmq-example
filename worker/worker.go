package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Create a new RabbitMQ connection with default settings.
	connRabbitMQ, err := amqp.Dial("amqp://rabbitmq:mypassword@localhost:5672/")
	if err != nil {
		panic(err)
	}

	// Open a new channel.
	channel, err := connRabbitMQ.Channel()
	if err != nil {
		log.Println(err)
	}
	defer channel.Close()

	// Start delivering queued messages.
	messages, err := channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	// Welcome message.
	log.Println("Successfully connected to RabbitMQ instance")
	log.Println("[*] - Waiting for messages")
	log.Println("[*] - Run Fiber API server and go to http://127.0.0.1:3000/send?msg=<YOUR TEXT HERE>")

	// Open a channel to receive messages.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			// For example, just show received message in console.
			log.Printf("Received message: %s\n", message.Body)
		}
	}()

	// Close the channel.
	<-forever
}
