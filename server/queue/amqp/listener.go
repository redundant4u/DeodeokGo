package amqp

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redundant4u/DeoDeokGo/queue"
	"github.com/redundant4u/DeoDeokGo/queue/contracts"
)

type amqpEventListener struct {
	connection *amqp.Connection
	exchange   string
	queue      string
}

func NewEventListener(conn *amqp.Connection, exchange string, queue string) queue.EventListener {
	listener := &amqpEventListener{
		connection: conn,
		exchange:   exchange,
		queue:      queue,
	}

	err := listener.setup()
	if err != nil {
		panic(err)
	}

	return listener
}

func (a *amqpEventListener) Listen(eventNames ...string) (<-chan queue.Event, <-chan error, error) {
	ch, err := a.connection.Channel()
	if err != nil {
		return nil, nil, err
	}

	defer ch.Close()

	for _, eventName := range eventNames {
		if err := ch.QueueBind(a.queue, eventName, a.exchange, false, nil); err != nil {
			return nil, nil, err
		}
	}

	msgs, err := ch.Consume(a.queue, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	events := make(chan queue.Event)
	errors := make(chan error)

	go func() {
		for msg := range msgs {
			rawEventName, ok := msg.Headers["x-event-name"]
			if !ok {
				errors <- fmt.Errorf("Msg did not contain x-event-name header")
				continue
			}

			eventName, ok := rawEventName.(string)
			if !ok {
				errors <- fmt.Errorf("x-event-name header is not string, but %t", rawEventName)
				continue
			}

			var event queue.Event

			switch eventName {
			case "eventCreated":
				event = new(contracts.EventCreatedEvent)
			default:
				errors <- fmt.Errorf("Event type %s is unknown", eventName)
				continue
			}

			err := json.Unmarshal(msg.Body, event)
			if err != nil {
				errors <- err
				continue
			}

			events <- event
		}
	}()

	return events, errors, nil
}

func (a *amqpEventListener) setup() error {
	ch, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(a.exchange, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(a.queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	return nil
}
