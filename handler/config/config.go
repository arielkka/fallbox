package config

import (
	"github.com/arielkka/rabbitmq"
	"github.com/spf13/viper"
)

type Config struct {
	Router  `mapstructure:"router"`
	Service `mapstructure:"service"`
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

	GetAllUserPNG string `mapstructure:"get_all_user_png"`
	GetUserPNG    string `mapstructure:"get_user_png"`
	PostUserPNG   string `mapstructure:"post_user_png"`
	DeleteUserPNG string `mapstructure:"delete_user_png"`

	GetAllUserJPG string `mapstructure:"get_all_user_jpg"`
	GetUserJPG    string `mapstructure:"get_user_jpg"`
	PostUserJPG   string `mapstructure:"post_user_jpg"`
	DeleteUserJPG string `mapstructure:"delete_user_jpg"`
}

type Message struct {
	DocumentPNGSend string `mapstructure:"document_png_send"`
	DocumentPNGGet  string `mapstructure:"document_png_get"`
	DocumentJPGSend string `mapstructure:"document_jpg_send"`
	DocumentEPGGet  string `mapstructure:"document_jpg_get"`
}

func NewServiceConfig(configFile string) (*Config, error) {
	var c *Config
	viper.SetConfigName(configFile)
	viper.AddConfigPath(".")
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

func NewRabbitMQConfig(configFile string) (*rabbitmq.Config, error) {
	var c *rabbitmq.Config
	viper.SetConfigName(configFile)
	viper.AddConfigPath(".")
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
