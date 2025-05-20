package mq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQPOptions struct {
	Username string
	Password string
	Host     string
	Port     int
	VHost    string
}

func NewAMQP(opt *AMQPOptions) *amqp.Connection {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		opt.Username,
		opt.Password,
		opt.Host,
		opt.Port,
		opt.VHost,
	)

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("failed to connect to rabbitmq: %v: url: %s", err, url)
	}

	return conn
}
