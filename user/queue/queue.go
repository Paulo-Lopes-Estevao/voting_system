package queue

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect() *amqp.Channel {

	dsn := "amqp://admin:admin@localhost:5672/"
	conn, err := amqp.Dial(dsn)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func UserMarkVote(payload []byte, exchange, routingKey, ExchangeKind string, ch *amqp.Channel) {

	err := ch.ExchangeDeclare(
		exchange,     // name
		ExchangeKind, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"vote_queue",
		true,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,     // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	err = ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
		})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Message processed")
}
