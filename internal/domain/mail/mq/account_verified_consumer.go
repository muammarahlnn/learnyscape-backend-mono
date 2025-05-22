package mq

import (
	"context"
	"fmt"
	authconstant "learnyscape-backend-mono/internal/domain/auth/constant"
	"learnyscape-backend-mono/internal/domain/auth/dto"
	"learnyscape-backend-mono/internal/domain/mail/constant"
	"learnyscape-backend-mono/internal/log"
	"learnyscape-backend-mono/pkg/mq"
	smtputil "learnyscape-backend-mono/pkg/util/smtp"
	"math"
	"sync"
	"time"

	"github.com/bytedance/sonic"
	amqp "github.com/rabbitmq/amqp091-go"
)

type AccountVerifiedConsumer struct {
	channel *amqp.Channel
	mailer  smtputil.Mailer
	queue   string
	wg      *sync.WaitGroup
}

func NewAccountVerifiedConsumer(conn *amqp.Connection, mailer smtputil.Mailer) mq.AMQPConsumer {
	queue := authconstant.AccountVerifiedQueue
	key := authconstant.AccountVerifiedKey
	exchange := authconstant.AccountVerifiedExchange

	ch, err := conn.Channel()
	if err != nil {
		log.Logger.Fatalf("failed to open a channel: %v", err)
	}

	if _, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		log.Logger.Fatalf("failed to declare a queue: %v", err)
	}

	if err := ch.QueueBind(
		queue,
		key,
		exchange,
		false,
		nil,
	); err != nil {
		log.Logger.Fatalf("failed to bind a queue: %v", err)
	}

	return &AccountVerifiedConsumer{
		channel: ch,
		mailer:  mailer,
		queue:   queue,
		wg:      &sync.WaitGroup{},
	}
}

func (c *AccountVerifiedConsumer) Consume(ctx context.Context, nWorker int) error {
	for i := 1; i <= nWorker; i++ {
		c.wg.Add(1)
		go c.Start(ctx, i)
	}

	return nil
}

func (c *AccountVerifiedConsumer) Start(ctx context.Context, workerID int) {
	defer c.wg.Done()

	msgs, err := c.channel.ConsumeWithContext(
		ctx,
		c.Queue(),
		fmt.Sprintf("%s-%d", c.Queue(), workerID),
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Logger.Errorf("failed to register %s-%d: %s", c.Queue(), workerID, err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Logger.Infof("%s-%d: shutting down...", c.Queue(), workerID)
			return
		case msg, ok := <-msgs:
			if !ok {
				log.Logger.Infof("%s-%d: message channel closed", c.Queue(), workerID)
				return
			}

			log.Logger.Infof("%v-%v: received a message %v", c.Queue(), workerID, string(msg.Body))
			for i := 1; i <= constant.AMQPRetryLimit+1; i++ {
				if err := c.Handler()(ctx, msg.Body); err != nil {
					log.Logger.Errorf("failed to consume message: %s", err)

					if i > constant.AMQPRetryLimit {
						log.Logger.Errorf("failed to consume message after %d retries: %s", constant.AMQPRetryLimit, err)
					} else {
						delay := math.Pow(constant.AMQPRetryDelay, float64(i))
						time.Sleep(time.Duration(delay) * time.Second)
						log.Logger.Infof("retrying to consume message, attempt %d", i)
					}
				} else {
					msg.Ack(false)
					break
				}
			}
		}
	}
}

func (c *AccountVerifiedConsumer) Handler() mq.AMQPHandler {
	return func(ctx context.Context, body []byte) error {
		var event dto.AccountVerifiedEvent
		if err := sonic.Unmarshal(body, &event); err != nil {
			log.Logger.Errorf("failed to unmarshal event: %s", err)
			return err
		}

		return c.mailer.SendMail(
			ctx,
			event.Email,
			constant.AccountVerifiedSubject,
			fmt.Sprintf(constant.AccountVerifiedTemplate, event.Name),
		)
	}
}

func (c *AccountVerifiedConsumer) Queue() string {
	return c.queue
}

func (c *AccountVerifiedConsumer) Close() error {
	log.Logger.Infof("Closing consumer for queue: %s", c.Queue())
	c.wg.Wait()
	return c.channel.Close()
}
