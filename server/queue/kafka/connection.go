package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

func NewKafkaClient(cfg *viper.Viper) sarama.Client {
	config := sarama.NewConfig()
	brokers := []string{cfg.GetString("broker.kafka.uri")}
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		panic("Colud not make kafka client: " + err.Error())
	}

	config.Producer.Return.Successes = true

	return client
}
