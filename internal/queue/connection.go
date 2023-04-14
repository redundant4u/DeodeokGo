package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func NewAmqpConnection(cfg *viper.Viper) *amqp.Connection {
	connection, err := amqp.Dial(cfg.GetString("amqp.uri"))

	if err != nil {
		panic("Could not establish AMQP connection: " + err.Error())
	}

	return connection
}
