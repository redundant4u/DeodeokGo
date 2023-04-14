package queue

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type EventEmitter interface {
	Emit(e Event) error
}

type amqpEventEmitter struct {
	context    context.Context
	connection *amqp.Connection
}

func NewEventEmitter(ctx context.Context, conn *amqp.Connection, exchange string) EventEmitter {
	emitter := &amqpEventEmitter{
		context:    ctx,
		connection: conn,
	}

	err := emitter.setup()
	if err != nil {
		panic(err)
	}

	return emitter
}

func (a *amqpEventEmitter) Emit(e Event) error {
	json, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
		return err
	}

	ch, err := a.connection.Channel()
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer ch.Close()

	msg := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": e.EventName()},
		ContentType: "application/json",
		Body:        json,
	}

	err = ch.PublishWithContext(a.context, "events", e.EventName(), false, false, msg)

	return err
}

func (a *amqpEventEmitter) setup() error {
	ch, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	err = ch.ExchangeDeclare("events", "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}

	return nil
}
