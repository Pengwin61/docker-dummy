package rabbitmq

import (
	"context"
	"docker-dummy/internal/config"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ClientRabbit struct {
	Client  *amqp.Connection
	Queue   *amqp.Queue
	Channel *amqp.Channel
}

func NewRabbitClient(config *config.Config) *ClientRabbit {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.RabbitMQ.User, config.RabbitMQ.Pass, config.RabbitMQ.Host, config.RabbitMQ.Port))
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
	}
	ch, q := newQueue(conn, config.RabbitMQ.QueueName)
	return &ClientRabbit{Client: conn, Queue: &q, Channel: ch}
}

func newQueue(conn *amqp.Connection, nameQueue string) (*amqp.Channel, amqp.Queue) {

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel. Error: %s", err)
	}

	q, err := ch.QueueDeclare(
		nameQueue, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}
	return ch, q
}

func (cr *ClientRabbit) Get(queue string) <-chan amqp091.Delivery {

	messages, err := cr.Channel.Consume(
		queue, // queue name
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // arguments
	)
	if err != nil {
		log.Println(err)
	}
	return messages
}

func (cr *ClientRabbit) Send(body string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := cr.Channel.PublishWithContext(ctx,
		"",            // exchange
		cr.Queue.Name, // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("failed to publish a message. Error: %s", err)
	}

	return body
}
