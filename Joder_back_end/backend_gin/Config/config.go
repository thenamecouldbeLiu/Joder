package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Mode string `mapstructure:"MODE"`
	Port string `mapstructure:"PORT"`
}

var config Config

func Init() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %v ", err))
	}

	if err := viper.Unmarshal(); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	log.WithFields(log.Fields{
		"val": config,
	}).Info("config loaded")
}
