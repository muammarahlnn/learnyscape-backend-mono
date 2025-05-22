package mq

import (
	"context"
	"fmt"
	"learnyscape-backend-mono/internal/domain/auth/constant"
	. "learnyscape-backend-mono/internal/domain/auth/constant"
	"learnyscape-backend-mono/internal/log"
	. "learnyscape-backend-mono/pkg/constant"
	"learnyscape-backend-mono/pkg/mq"
	"math"
	"time"

	"github.com/bytedance/sonic"
	amqp "github.com/rabbitmq/amqp091-go"
)

type SendVerificationPublisher struct {
	Channel  *amqp.Channel
	exchange string
}

func NewSendVerificationPublisher(conn *amqp.Connection) mq.AMQPPublisher {
	exchange := SendVerificationExchange
	ch, err := conn.Channel()
	if err != nil {
		log.Logger.Fatalf("failed to open a channel: %s", err)
	}

	if err := ch.ExchangeDeclare(exchange, ExchangeTypeTopic, true, false, false, false, nil); err != nil {
		if amqpErr, ok := err.(*amqp.Error); ok && amqpErr.Code != amqp.PreconditionFailed {
			log.Logger.Fatalf("failed to declare an exchange: %s", err)
		}
	}

	return &SendVerificationPublisher{
		Channel:  ch,
		exchange: exchange,
	}
}

func (p *SendVerificationPublisher) Publish(ctx context.Context, event mq.AMQPEvent) error {
	bytes, err := sonic.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	for i := 1; i <= AMQPRetryLimit; i++ {
		err = p.Channel.PublishWithContext(
			ctx,
			p.exchange,
			event.Key(),
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        bytes,
			},
		)
		if err == nil {
			log.Logger.Infof("message published: %s", string(bytes))
			return nil
		}

		log.Logger.Errorf("failed to publish message: %s", err)

		delay := time.Duration(math.Pow(2, float64(i))) * constant.AMQPRetryDelay * time.Second
		time.Sleep(delay)
	}

	log.Logger.Errorf("message is aborted: %s", err)
	return err
}

func (p *SendVerificationPublisher) Exchange() string {
	return p.exchange
}
