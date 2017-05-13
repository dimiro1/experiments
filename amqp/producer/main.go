package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"messages",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello World"),
	})
	if err != nil {
		log.Fatal(err)
	}
}
