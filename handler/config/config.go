package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Router
	Service
}

type Service struct {
	Name string
	Message
}

type Router struct {
	Host          string
	Port          string
	Cookie        string
	AuthPath      string
	PngSendPath   string
	PngGetPath    string
	excelSendPath string
	excelGetPath  string
}

type Message struct {
	DocumentPngSend   string
	DocumentPngGet    string
	DocumentExcelSend string
	DocumentExcelGet  string
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
