package config

import (
	"github.com/spf13/viper"
)

var config *viper.Viper

func LoadConfig(env string) (*viper.Viper, error) {
	config = viper.New()

	config.SetConfigName(env)
	config.SetConfigType("yaml")
	config.AddConfigPath("config/")

	err := config.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return config, nil
}

func GetConfig() *viper.Viper {
	return config
}
