package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func LoadConfig(env string) *viper.Viper {
	config = viper.New()

	config.SetConfigName(env)
	config.SetConfigType("yaml")
	config.AddConfigPath("config/")

	err := config.ReadInConfig()

	if err != nil {
		log.Fatal(fmt.Errorf("Fatal error config file: %w", err))
	}

	return config
}

func GetConfig() *viper.Viper {
	return config
}
