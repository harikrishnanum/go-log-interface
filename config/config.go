package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Logger string `json:"logger"`
}

var Conf *Config

func init() {
	Conf = &Config{}
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
