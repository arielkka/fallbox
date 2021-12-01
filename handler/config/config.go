package config

import (
	"github.com/arielkka/rabbitmq"
	"github.com/spf13/viper"
)

type Config struct {
	Router   `mapstructure:"router"`
	Service  `mapstructure:"service"`
	Database `mapstructure:"db"`
}

type Service struct {
	Name    string `mapstructure:"name"`
	Message `mapstructure:"message"`
}

type Router struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	Cookie string `mapstructure:"cookie"`

	AuthPath     string `mapstructure:"auth_path"`
	Registration string `mapstructure:"registration_path"`

	GetUserExcel    string `mapstructure:"get_excel_png"`
	PostUserExcel   string `mapstructure:"post_excel_png"`
	DeleteUserExcel string `mapstructure:"delete_excel_png"`

	GetUserTxt    string `mapstructure:"get_txt_jpg"`
	PostUserTxt   string `mapstructure:"post_txt_jpg"`
	DeleteUserTxt string `mapstructure:"delete_txt_jpg"`
}

type Message struct {
	DocumentExcelSend   string `mapstructure:"document_excel_send"`
	DocumentExcelGet    string `mapstructure:"document_excel_get"`
	DocumentExcelDelete string `mapstructure:"document_excel_delete"`

	DocumentTXTSend   string `mapstructure:"document_txt_send"`
	DocumentTXTGet    string `mapstructure:"document_txt_get"`
	DocumentTXTDelete string `mapstructure:"document_txt_delete"`
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
