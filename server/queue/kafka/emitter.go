package kafka

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/redundant4u/DeoDeokGo/queue"
)

type kafkaEventEmitter struct {
	producer sarama.SyncProducer
}

type messageEnvelope struct {
	EventName string      `json:"eventName"`
	Payload   interface{} `json:"payload"`
}

func NewKafkaEventEmitter(client sarama.Client) queue.EventEmitter {
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}

	emitter := &kafkaEventEmitter{
		producer: producer,
	}

	return emitter
}

func (k *kafkaEventEmitter) Emit(e queue.Event) error {
	envelop := messageEnvelope{e.EventName(), e}
	jsonBody, err := json.Marshal(&envelop)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: e.EventName(),
		Value: sarama.ByteEncoder(jsonBody),
	}

	_, _, err = k.producer.SendMessage(msg)

	return err
}
