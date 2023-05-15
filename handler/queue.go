package handler

import (
	"gbEngine/admin"
	"gbEngine/config"

	"github.com/streadway/amqp"
)

type QueueHandler struct {
	Queue  config.Queue
	Logger *admin.Logger
}

func (queue *QueueHandler) Produce(queueName string, data []byte) error {
	err := queue.Queue.Channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	return err
}
