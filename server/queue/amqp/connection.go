package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func NewAmqpConnection(cfg *viper.Viper) (*amqp.Connection, error) {
	conn, err := amqp.Dial(cfg.GetString("broker.amqp.uri"))

	if err != nil {
		return nil, err
	}

	return conn, nil
}
