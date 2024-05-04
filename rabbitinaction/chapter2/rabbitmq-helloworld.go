package main

import (
	"fmt"
	"time"
	"log"
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}

func main() {

	credentials := amqp.PlainAuth{Username: "guest", Password: "guest"}
	config := amqp.Config{
		SASL: []amqp.Authentication{
			&credentials,
		},
	}

	// Connecting to the rabbitmq server
	conn, err := amqp.DialConfig("amqp://localhost:5672/", config)
	failOnError(err, "something went wrong")
	defer conn.Close()

	// open a channel
	ch, err := conn.Channel()
	failOnError(err, "failed to open channel...")
	defer ch.Close()

	//Set confirm mode
	err = ch.Confirm(false)
	failOnError(err, "unable to confirm")

	// Create the exchange
	err = ch.ExchangeDeclare(
		"kato-exchange",
		"direct",
		true,
		true,
		false,
		true,
		nil,
	)
	failOnError(err, "Unable to Declare exchange")
	fmt.Println("Exchange Created successfully")

	//Create the queue
	q, err := ch.QueueDeclare(
		"kato-queue",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "queue did not get created")
	log.Println(q.Name)

	err = ch.QueueBind(
		"kato-queue",
		"",
		"kato-exchange",
		false,
		nil,
	)
	failOnError(err, "Unable to bind queue to change")
	log.Println("Queue bind successful")

	// Provisioning the message
	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp: time.Now(),
		ContentType: "text/plain",
		Body: []byte("Hello Kato World!"),
	}

	// Publishing a message
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = ch.PublishWithContext(
		ctx,
		"kato-exchang",
	  "",
		true,
		false,
		msg,
	)
	failOnError(err, "unable to publish message...")

	// confirm message delibery
	confirm := <-ch.NotifyPublish(make(chan amqp.Confirmation, 1))
	if confirm.Ack {
		log.Println("Message confirmed!")
	} else {
		log.Println("Message Failed!")
	}


	msgs, err := ch.Consume(
		"kato-queue",
		"kato-consumer",
		true,
		false, //make sure it is the sole consumer of the queue
		false,
		false,
		nil,
	)
	failOnError(err, "Unable to consume message")

	for d := range msgs {
		log.Printf("Got a message: %s", d.Body)
		break
	}
	

}
