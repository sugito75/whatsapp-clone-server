package mq

import (
	"encoding/json"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"
)

type messageQueue struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
}

func New(conn *amqp091.Connection) *messageQueue {
	channel, err := conn.Channel()
	if err != nil {
		slog.Error(err.Error())
		conn.Close()
		return nil
	}

	return &messageQueue{
		conn:    conn,
		channel: channel,
	}
}

func (mq *messageQueue) Publish(q string, m Message) error {
	body, _ := json.Marshal(m)

	return mq.channel.Publish(
		"",
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

}

func (mq *messageQueue) Consume(q string) (<-chan amqp091.Delivery, error) {
	return mq.channel.Consume(q, "", true, false, false, false, nil)
}
