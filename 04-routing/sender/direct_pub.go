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
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	helper.FailOnError(err, "Failed to declare an exhange")

	body := helper.BodyFrom(os.Args)
	err = ch.Publish(
		"logs_direct",                // exchange
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

// go run direct_pub.go info "run it info"
// go run direct_pub.go warning "run it warn"
// go run direct_pub.go error "run it error"
