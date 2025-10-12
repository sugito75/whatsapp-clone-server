package config

import (
	"fmt"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

func GetRabbitMQConn() *amqp091.Connection {
	conn, err := amqp091.Dial(os.Getenv("RABBIT_MQ_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)

	}

	return conn
}
