package rabbitmq

import (
	"github.com/arielkka/rabbitmq"
)

func NewClient(serviceName string, cfg *rabbitmq.Config) (*rabbitmq.Client, error){
	client, err := rabbitmq.NewClient(serviceName, cfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}
