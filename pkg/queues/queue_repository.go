package queue

import (
	"github.com/ivanmeca/DesafioPismo/v2/config"
	"github.com/streadway/amqp"
	"strconv"
)

type BindParams struct {
	Name     string
	Key      string
	Exchange string
	NoWait   bool
	Args     map[string]interface{}
}

type Repository struct {
	connection *amqp.Connection
}

func NewQueueRepository(conf config.QueueConfig) (*Repository, error) {

	auth := amqp.PlainAuth{Username: conf.Login, Password: conf.Password}
	var arrAuth []amqp.Authentication
	arrAuth = append(arrAuth, &auth)
	config := amqp.Config{
		SASL: arrAuth,
	}
	conn, err := amqp.DialConfig("amqp://"+conf.Host+":"+strconv.Itoa(conf.Port)+"/", config)
	if err != nil {
		return nil, err
	}

	queueRp := Repository{}
	queueRp.connection = conn
	return &queueRp, nil
}

func NewQueueBindParams(name, key, exchange string) BindParams {
	return BindParams{
		Name:     name,
		Key:      key,
		Exchange: exchange,
		NoWait:   false,
		Args:     nil,
	}
}

func (q *Repository) QueueBind(params BindParams) error {
	activeChannel, err := q.connection.Channel()
	if err != nil {
		return err
	}
	if err := activeChannel.QueueBind(
		params.Name,
		params.Key,
		params.Exchange,
		params.NoWait,
		params.Args); err != nil {
		return err
	}
	return nil
}

func (q *Repository) QueueDeclare(params Params, withErrorQueue bool) (*Queue, error) {
	if withErrorQueue {
		q.errorQueueDeclare(params)
	}
	return q.queueDeclare(params)
}

func (q *Repository) ExchangeDeclare(params Params) error {
	activeChannel, err := q.connection.Channel()
	if err != nil {
		return err
	}
	if err := activeChannel.ExchangeDeclare(
		params.Name,
		params.Kind,
		params.Durable,
		params.AutoDelete,
		params.Internal,
		params.NoWait,
		params.Args); err != nil {
		return err
	}
	return nil
}

func (q *Repository) errorQueueDeclare(params Params) error {

	errorQueueName := params.Name + "-error"

	exchangeParams := NewQueueParams("Error")
	q.ExchangeDeclare(exchangeParams)

	args := map[string]interface{}{
		"x-dead-letter-exchange":    "Error",
		"x-dead-letter-routing-key": errorQueueName,
	}
	qParam := NewQueueParams(errorQueueName)
	qParam.Args = args
	_, err := q.queueDeclare(qParam)

	if err != nil {
		return err
	}

	err = q.QueueBind(NewQueueBindParams(errorQueueName, errorQueueName, "Error"))
	if err != nil {
		return err
	}
	return nil
}

func (q *Repository) queueDeclare(params Params) (*Queue, error) {
	queue, err := NewQueue(params, q.connection)
	if err != nil {
		return nil, err
	}
	return queue, nil
}
