package main

import (
	"log"
	"github.com/streadway/amqp"
)

// Parse error if catch.
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Make a connection to CloudAMQP.
	conn, err := amqp.Dial("amqp://user:password@crane.rmq.cloudamqp.com/vhost")
	failOnError(err, "Failed to connect to CloudAMQP")
	defer conn.Close();

	// Create a channel.
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close();

	// Queue name must be the same with publisher
	queueName := "profile";
	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	// Listen to the queue.
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	// Make a go routine by using anonymous function
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	// Always listening for incoming message from message broker.
	<- forever
}




