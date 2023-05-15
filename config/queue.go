package config

import "github.com/streadway/amqp"

type Queue struct {
	queue   *amqp.Queue
	Channel *amqp.Channel
	Jobs    <-chan amqp.Delivery
}

func ConnectToQueue(env *ENV) (*Queue, error) {
	var address string = "amqp://" + env.RabbitMQUsername + ":" + env.RabbitMQPassword + "@" + env.RabbitMQHost + ":" + env.RabbitMQPort + "/"
	conn, err := amqp.Dial(address)
	if err != nil {
		return nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	queue, err := channel.QueueDeclare(env.EngineName, false, true, false, false, nil)
	if err != nil {
		return nil, err
	}

	jobsChan, err := channel.Consume(env.EngineName, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &Queue{
		queue:   &queue,
		Channel: channel,
		Jobs:    jobsChan,
	}, nil
}
