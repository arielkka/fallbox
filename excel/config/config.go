package config

import (
	"github.com/arielkka/rabbitmq"
	"github.com/spf13/viper"
)

type Config struct {
	Service  `mapstructure:"service"`
	Database `mapstructure:"db"`
}

type Service struct {
	Name    string `mapstructure:"name"`
	Message `mapstructure:"message"`
}

type Message struct {
	DocumentExcelSend   string `mapstructure:"document_excel_send"`
	DocumentExcelGet    string `mapstructure:"document_excel_get"`
	DocumentExcelDelete string `mapstructure:"document_excel_delete"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

func NewServiceConfig(path string) (*Config, error) {
	var c *Config
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewRabbitMQConfig(path string) (*rabbitmq.Config, error) {
	var c *rabbitmq.Config
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
