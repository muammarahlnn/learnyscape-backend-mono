package mq

import "context"

type AMQPPublisher interface {
	Publish(ctx context.Context, event AMQPEvent) error
	Exchange() string
}

type AMQPEvent interface {
	Key() string
}
