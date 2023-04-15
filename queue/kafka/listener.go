package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/mitchellh/mapstructure"
	"github.com/redundant4u/DeoDeokGo/queue"
	"github.com/redundant4u/DeoDeokGo/queue/contracts"
)

type kafkaEventListener struct {
	consumer   sarama.Consumer
	partitions []int32
}

func NewKafkaEventListener(client sarama.Client, partitions []int32) queue.EventListener {
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}

	listener := &kafkaEventListener{
		consumer:   consumer,
		partitions: partitions,
	}

	return listener
}

func (k *kafkaEventListener) Listen(events ...string) (<-chan queue.Event, <-chan error, error) {
	var err error

	topic := "events"
	results := make(chan queue.Event)
	errors := make(chan error)

	partitions := k.partitions
	if len(partitions) == 0 {
		partitions, err = k.consumer.Partitions(topic)
		if err != nil {
			return nil, nil, err
		}
	}

	log.Printf("topic %s has partitions: %v\n", topic, partitions)

	for _, partition := range partitions {
		con, err := k.consumer.ConsumePartition(topic, partition, 0)
		if err != nil {
			return nil, nil, err
		}

		go func() {
			for msg := range con.Messages() {
				body := messageEnvelope{}
				err := json.Unmarshal(msg.Value, &body)
				if err != nil {
					errors <- fmt.Errorf("Could not JSON-decode message: %s", err)
					continue
				}

				var event queue.Event

				switch body.EventName {
				case "event.created":
					event = &contracts.EventCreatedEvent{}
				default:
					errors <- fmt.Errorf("Unknown event type : %s", body.EventName)
					continue
				}

				cfg := mapstructure.DecoderConfig{
					Result:  event,
					TagName: "json",
				}
				decoder, _ := mapstructure.NewDecoder(&cfg)
				err = decoder.Decode(body.Payload)
				if err != nil {
					errors <- fmt.Errorf("Colud not map event %s: %s", body.EventName, err)
				}

				results <- event
			}
		}()
	}

	return results, errors, nil
}
