package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Mode               string `mapstructure:"MODE"`
	Port               string `mapstructure:"PORT"`
	GoogleSecretKey    string `mapstructure:"GOOGLE_SECRET_KEY"`
	GoogleClientId     string `mapstructure:"GOOLE_CLIENT_ID"`
	GoogleScope		   string `mapstructure:"GOOGLE_SCOPE"`
	GoogleRedirect	   string `mapstructure:"GOOGLE_REDIRECT"`
	JWT_TOKEN_LIFE     int64  `mapstructure:"JWT_TOKEN_LIFE"`
	JWT_KEY            string `mapstructure:"JWT_KEY"`
	COOKIE_KEY         string `mapstructure:"COOKIE_KEY"`
	MYSQL_ACCNT        string `mapstructure:"MYSQL_ACCNT"`
	MYSQL_PASSWORD     string `mapstructure:"MYSQL_PASSWORD"`
	MYSQL_URL_IP       string `mapstructure:"MYSQL_URL_IP"`
	MY_SQL_SCHEMA      string `mapstructure:"MY_SQL_SCHEMA"`
	MY_SQL_CHARSET     string `mapstructure:"MY_SQL_CHARSET"`
	MY_SQL_MAXLIFETIME int    `mapstructure:"MY_SQL_MAXLIFETIME"`
	MY_SQL_MAXCONN     int    `mapstructure:"MY_SQL_MAXCONN"`
	MYSQL_MAXIDEL      int    `mapstructure:"MYSQL_MAXIDEL"`
}

var Val Config

func init() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %v ", err))
	}

	if err := viper.Unmarshal(&Val); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	log.WithFields(log.Fields{
		"val": Val,
	}).Info("config loaded")
}
