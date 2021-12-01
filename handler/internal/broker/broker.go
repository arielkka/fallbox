package broker

import (
	"github.com/arielkka/fallbox/handler/config"
	"github.com/arielkka/rabbitmq"
	"github.com/streadway/amqp"
	"time"
)

type Broker struct {
	client   *rabbitmq.Client
	cfg      *config.Config
	messages map[string]chan *rabbitmq.Message
}

func NewBroker(client *rabbitmq.Client, cfg *config.Config) *Broker {
	return &Broker{client: client, cfg: cfg, messages: InitDelivery(cfg)}
}

func InitDelivery(cfg *config.Config) map[string]chan *rabbitmq.Message {
	m := make(map[string]chan *rabbitmq.Message, 0)
	m[cfg.Service.Message.DocumentPNGGet] = make(chan *rabbitmq.Message)
	m[cfg.Service.Message.DocumentPNGSend] = make(chan *rabbitmq.Message)
	m[cfg.Service.Message.DocumentPNGDelete] = make(chan *rabbitmq.Message)
	m[cfg.Service.Message.DocumentJPGGet] = make(chan *rabbitmq.Message)
	m[cfg.Service.Message.DocumentJPGSend] = make(chan *rabbitmq.Message)
	m[cfg.Service.Message.DocumentJPGDelete] = make(chan *rabbitmq.Message)
	return m
}

func (b *Broker) Ping() error {
	return b.client.Ping()
}

func (b *Broker) Publish(messageType string, correlationID string, body []byte) error {
	err := b.client.SendMessage(messageType, correlationID, body)
	if err != nil {
		return err
	}
	return nil
}

func (b *Broker) Subscribe(messageType string, correlationID string) ([]byte, error) {
	for message := range b.messages[messageType] {
		if message.GetID() == correlationID {
			return message.GetBody(), nil
		}
	}
	return nil, amqp.Error{Code: 400}
}

func (b *Broker) CreateListeners() error {
	errors := make(chan error)
	go func() {
		err := b.client.StartConsumer(
			b.cfg.Service.Message.DocumentPNGGet,
			b.messages[b.cfg.Service.Message.DocumentPNGGet],
			true,
		)
		if err != nil {
			errors <- err
		}
	}()

	go func() {
		err := b.client.StartConsumer(
			b.cfg.Service.Message.DocumentJPGGet,
			b.messages[b.cfg.Service.Message.DocumentJPGGet],
			true)
		if err != nil {
			errors <- err
		}
	}()

	go func() {
		err := b.client.StartConsumer(
			b.cfg.Service.Message.DocumentJPGGet,
			b.messages[b.cfg.Service.Message.DocumentPNGDelete],
			true)
		if err != nil {
			errors <- err
		}
	}()


	go func() {
		err := b.client.StartConsumer(
			b.cfg.Service.Message.DocumentJPGGet,
			b.messages[b.cfg.Service.Message.DocumentJPGDelete],
			true)
		if err != nil {
			errors <- err
		}
	}()

	const delay = 10
	select {
	case err := <-errors:
		return err
	case <-time.After(delay * time.Millisecond):
		return nil
	}
}

func (b *Broker) CreateExchanges() error {
	err := b.client.CreateExchange(b.cfg.Service.Message.DocumentJPGSend)
	if err != nil {
		return err
	}
	err = b.client.CreateExchange(b.cfg.Service.Message.DocumentJPGGet)
	if err != nil {
		return err
	}
	err = b.client.CreateExchange(b.cfg.Service.Message.DocumentPNGGet)
	if err != nil {
		return err
	}
	err = b.client.CreateExchange(b.cfg.Service.Message.DocumentPNGSend)
	if err != nil {
		return err
	}
	err = b.client.CreateExchange(b.cfg.Service.Message.DocumentPNGDelete)
	if err != nil {
		return err
	}
	err = b.client.CreateExchange(b.cfg.Service.Message.DocumentJPGDelete)
	if err != nil {
		return err
	}
	return nil
}

func (b *Broker) CreateQueues() error {
	err := b.client.CreateQueue(b.cfg.Service.Message.DocumentPNGGet, true)
	if err != nil {
		return err
	}
	err = b.client.CreateQueue(b.cfg.Service.Message.DocumentJPGGet, true)
	if err != nil {
		return err
	}
	err = b.client.CreateQueue(b.cfg.Service.Message.DocumentPNGDelete, true)
	if err != nil {
		return err
	}
	err = b.client.CreateQueue(b.cfg.Service.Message.DocumentJPGDelete, true)
	if err != nil {
		return err
	}
	return nil
}
