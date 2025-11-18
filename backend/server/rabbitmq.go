package server

import (
	"log"
	"github.com/rabbitmq/amqp091-go"
)

type MQ struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
	Queue   amqp091.Queue
}

func NewRabbitMQ() *MQ {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Error conectando RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error creando canal: %v", err)
	}

	q, err := ch.QueueDeclare(
		"transmutaciones_queue",
		true,  
		false, 
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error creando queue: %v", err)
	}

	return &MQ{
		Conn:    conn,
		Channel: ch,
		Queue:   q,
	}
}
