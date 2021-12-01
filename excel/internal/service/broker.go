package service

type Broker interface {
	Publish(messageType string, correlationID string, body []byte) error
	Subscribe(messageType string, correlationID string) ([]byte, error)
	Ping() error
}
