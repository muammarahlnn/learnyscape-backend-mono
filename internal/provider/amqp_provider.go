package provider

import (
	mailmq "learnyscape-backend-mono/internal/domain/mail/mq"
	"learnyscape-backend-mono/pkg/mq"
)

func BootstrapAMQP() []mq.AMQPConsumer {
	return []mq.AMQPConsumer{
		mailmq.NewSendVerificationConsumer(rabbitmq, mailer),
	}
}
