package tools

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

type QueueProducer struct {
	amqpCh   *amqp.Channel
	message  *domain.Message
	exchange string
	routing  string
	queue    string
}

func NewQueueProducer(amqpCh *amqp.Channel, message *domain.Message, queueConfig *domain.QueueConfiguration) *QueueProducer {
	return &QueueProducer{
		amqpCh:   amqpCh,
		message:  message,
		exchange: queueConfig.Exchange,
		routing:  queueConfig.RoutingKey,
		queue:    queueConfig.Queue,
	}
}

func (q *QueueProducer) InitializeQueue() error {
	err := q.InitializeExchange()
	if err != nil {
		return err
	}
	queue, err := infra.DeclareQueue(q.amqpCh, q.queue)
	if err != nil {
		return err
	}
	err = infra.BindQueue(q.amqpCh, queue, q.exchange, q.routing)
	if err != nil {
		return err
	}
	return nil
}

func (q *QueueProducer) InitializeExchange() error {
	err := infra.DeclareExchange(q.amqpCh, q.exchange)
	if err != nil {
		return err
	}
	return nil
}

func (q *QueueProducer) Publish() error {
	messageBytes, err := json.Marshal(q.message)
	if err != nil {
		return err
	}
	err = q.amqpCh.Publish(
		q.exchange,
		q.routing,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        messageBytes,
		})
	if err != nil {
		return err
	}
	return nil
}
