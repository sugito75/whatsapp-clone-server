package mq

import "github.com/rabbitmq/amqp091-go"

type MessageQueue interface {
	Publish(m Message) error
	Consume(q string) (<-chan amqp091.Delivery, error)
}
