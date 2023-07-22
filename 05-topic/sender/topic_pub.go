package main

import (
	"log"
	"message-broker/helper"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://user:user@localhost:5672/")
	helper.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	helper.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// exhange is push messages to queues
	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	helper.FailOnError(err, "Failed to declare an exhange")

	body := helper.BodyFrom(os.Args)
	err = ch.Publish(
		"logs_topic",                 // exchange
		helper.SeverityFrom(os.Args), // routing key
		false,                        // mandatory
		false,                        //immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	helper.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

// go run topic_pub.go "kern.info" "An info kernel error"

// go run topic_pub.go "info.info" "An info kernel error"
