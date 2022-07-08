package queue

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/streadway/amqp"
	rabbitmq "github.com/wagslane/go-rabbitmq"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var consumerName = "example"

func ConsumerStart(ExchangeName, ExchangeKind, queue string, routingKeys []string, messages chan string) {
	consumer, err := rabbitmq.NewConsumer(
		"amqp://admin:admin@localhost:5672/", rabbitmq.Config{},
		rabbitmq.WithConsumerOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := consumer.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = consumer.StartConsuming(
		func(d rabbitmq.Delivery) rabbitmq.Action {
			log.Printf("consumed: %v", string(d.Body))
			go func() { messages <- "ping" }()
			// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
			return rabbitmq.Ack
		},
		queue,
		routingKeys,
		rabbitmq.WithConsumeOptionsConcurrency(1),
		rabbitmq.WithConsumeOptionsQueueDurable,
		rabbitmq.WithConsumeOptionsBindingExchangeName(ExchangeName),
		rabbitmq.WithConsumeOptionsBindingExchangeKind(ExchangeKind),
		rabbitmq.WithConsumeOptionsBindingExchangeDurable,
		rabbitmq.WithConsumeOptionsConsumerName(consumerName),
	)
	if err != nil {
		log.Fatal(err)
	}

	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("stopping consumer")
}

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

/*
exchange string -- exchange name

routingKey string -- routing key

ExchangeKind string -- types are "direct", "fanout", "topic" and "headers"

*/
func ShowVotes(exchange, routingKey, ExchangeKind string, ch *amqp.Channel, in chan []byte) {

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

	if err != nil {
		panic(err.Error())
	}

	if q.Messages <= 0 {
		in <- []byte("")
	}

	err = ch.QueueBind(
		q.Name,     // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name,
		"vote",
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		panic(err.Error())
	}

	go func() {
		for m := range msgs {
			in <- []byte(m.Body)
		}
		close(in)
	}()
}
