package treater

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"time"
)

type AmqpQueueProtocol struct {
	amqpCh            *amqp.Channel
	messageDispatcher MessageTreater
	queue             string
}

func NewAmqpQueueProtocol(amqpCh *amqp.Channel, queue string, messageTreater MessageTreater) *AmqpQueueProtocol {
	return &AmqpQueueProtocol{
		amqpCh:            amqpCh,
		messageDispatcher: messageTreater,
		queue:             queue,
	}
}

func (a *AmqpQueueProtocol) InitializeConsumer(ctx context.Context, loop bool, timeout int) error {
	logger := zap.Logger{}
	logger.Debug("initialize consumer", zap.String("queue name", a.queue))
	msgs, err := a.amqpCh.Consume(
		a.queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "error to try consume a queue")
	}
	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	if timeout < 0 {
		timer.Stop()
	}
	if loop {
		for {
			err, done := a.startConsume(ctx, msgs, timer)
			if done {
				return err
			}
		}
	}
	err, _ = a.startConsume(ctx, msgs, timer)
	return err
}

func (a *AmqpQueueProtocol) startConsume(ctx context.Context, msgs <-chan amqp.Delivery) (error, bool) {
	logger := zap.Logger{}
	logger.Debug("start consume")
	select {
	case <-ctx.Done():
		return nil, true
	case msg, ok := <-msgs:
		if !ok {
			return nil, true
		}
		a.handleMessage(ctx, msg)
	}
	return nil, false
}

func (a *AmqpQueueProtocol) handleMessage(ctx context.Context, amqpMsg amqp.Delivery) {
	logger := log.Log(ctx)
	body := amqpMsg.Body
	message, err := event.NewMessageBuilder().BuildMessage(body)
	if err != nil {
		logger.With(zap.Error(err)).Error("build message")
		err = a.publishOnErrorQueue(message)
		if err != nil {
			err = amqpMsg.Nack(true, true)
			if err != nil {
				logger.With(zap.Error(err)).Error("fail to nack message")
				return
			}
		}
		return
	}
	logger.Debug("message dispatcher start handle message")
	if a.messageDispatcher.HandleMessage(ctx, message) {
		err = amqpMsg.Ack(true)
		if err != nil {
			logger.With(zap.Error(err)).Error("ack message")
			err = a.publishOnErrorQueue(message)
			if err != nil {
				err = amqpMsg.Nack(true, true)
				if err != nil {
					logger.With(zap.Error(err)).Error("fail to nack message")
					return
				}
			}
		}
		return
	}
}

func (a *AmqpQueueProtocol) publishOnErrorQueue(msg *domain.Message) error {
	_, err := a.createQueue(msg.ErrorQueue)
	if err != nil {
		return err
	}
	amqpPub, err := a.getAmqpMessage(msg, "")
	if err != nil {
		return err
	}
	err = a.publishMessage(amqpPub, msg.ErrorQueue)
	if err != nil {
		return err
	}
	return nil
}

func (a *AmqpQueueProtocol) createQueue(queueName string) (string, error) {
	q, err := a.amqpCh.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return "", err
	}
	return q.Name, nil
}

func (a *AmqpQueueProtocol) publishMessage(amqpPub amqp.Publishing, errorQueue string) error {
	err := a.amqpCh.Publish(
		"",
		errorQueue,
		false,
		false,
		amqpPub,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *AmqpQueueProtocol) getAmqpMessage(msg *domain.Message, timeout string) (amqp.Publishing, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return amqp.Publishing{}, err
	}
	amqpPub := amqp.Publishing{
		ContentType:  "text/plain",
		Body:         body,
		DeliveryMode: 2,
	}
	if timeout != "" {
		amqpPub.Expiration = timeout
	}
	return amqpPub, nil
}
