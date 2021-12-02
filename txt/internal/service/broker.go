package service

import "github.com/arielkka/fallbox/txt/internal/models"

type Broker interface {
	Publish(messageType string, correlationID string, body []byte) error
	Subscribe(messageType string) (*models.Message, error)
	Respond(consumer, msg, corrID string, body []byte) error
	Ping() error
}
