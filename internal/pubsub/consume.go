package pubsub

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SimpleQueueType int

const (
	SimpleQueueTypeDurable SimpleQueueType = iota
	SimpleQueueTypeTransient
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	simpleQueueType SimpleQueueType,
) (*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, fmt.Errorf("could not declare channel: %v", err)
	}
	defer ch.Close()
	fmt.Println("Channel created successful")

	queue, err := ch.QueueDeclare(
		queueName,
		simpleQueueType == SimpleQueueTypeDurable,
		simpleQueueType == SimpleQueueTypeTransient,
		simpleQueueType == SimpleQueueTypeTransient,
		false,
		nil,
	)
	if err != nil {
		return nil, amqp.Queue{}, fmt.Errorf("could not declare queue: %v", err)
	}

	err = ch.QueueBind(queueName, key, exchange, false, nil)
	if err != nil {
		return nil, amqp.Queue{}, fmt.Errorf("could not bind queue: %v", err)
	}

	return ch, queue, nil
}
