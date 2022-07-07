package rmq

import (
	"fmt"
	"gbEngine/redisAction"
	"log"
	"math/rand"

	"github.com/streadway/amqp"
)

type RMQ struct {
	Name     string
	Address  string
	Port     string
	Username string
	Password string
	RedisDB  *redisAction.Redis
	Msgs     <-chan amqp.Delivery
	ch       *amqp.Channel
	q        *amqp.Queue
}

func (r *RMQ) Init() {
	var address string = "amqp://" + r.Username + ":" + r.Password + "@" + r.Address + ":" + r.Port + "/"
	conn, err := amqp.Dial(address)
	if err != nil {
		fmt.Println("[ERROR] : ", err.Error())
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("[ERROR] : ", err.Error())
	}
	r.ch = ch
	q, err := ch.QueueDeclare(
		r.Name,
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err.Error())
	}
	r.q = &q
	r.Msgs, err = r.ch.Consume(
		r.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("RMQ client Connected")
}

// func (r *RMQ) deleteChannel() {
// 	r.ch.
// }

func (r *RMQ) getEngineChannel() string {
	names := r.RedisDB.GetEngineName()
	randomIndex := rand.Intn(len(names))
	pick := names[randomIndex]
	return pick
}

func (r *RMQ) Produce(name string, job []byte) error {
	err := r.ch.Publish(
		"",
		name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        job,
		},
	)
	return err
}

func (r *RMQ) Consume() {
	var err error
	r.Msgs, err = r.ch.Consume(
		r.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err.Error())
	}
}
