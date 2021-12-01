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
	m[cfg.Service.Message.DocumentExcelGet] = make(chan *rabbitmq.Message)
	m[cfg.Service.Message.DocumentExcelSend] = make(chan *rabbitmq.Message)
	m[cfg.Service.Message.DocumentExcelDelete] = make(chan *rabbitmq.Message)
	m[cfg.Service.Message.DocumentTXTGet] = make(chan *rabbitmq.Message)
	m[cfg.Service.Message.DocumentTXTSend] = make(chan *rabbitmq.Message)
	m[cfg.Service.Message.DocumentTXTDelete] = make(chan *rabbitmq.Message)
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
			b.cfg.Service.Message.DocumentExcelGet,
			b.messages[b.cfg.Service.Message.DocumentExcelGet],
			true,
		)
		if err != nil {
			errors <- err
		}
	}()

	go func() {
		err := b.client.StartConsumer(
			b.cfg.Service.Message.DocumentTXTGet,
			b.messages[b.cfg.Service.Message.DocumentTXTGet],
			true)
		if err != nil {
			errors <- err
		}
	}()

	go func() {
		err := b.client.StartConsumer(
			b.cfg.Service.Message.DocumentExcelSend,
			b.messages[b.cfg.Service.Message.DocumentExcelSend],
			true,
		)
		if err != nil {
			errors <- err
		}
	}()

	go func() {
		err := b.client.StartConsumer(
			b.cfg.Service.Message.DocumentTXTSend,
			b.messages[b.cfg.Service.Message.DocumentTXTSend],
			true)
		if err != nil {
			errors <- err
		}
	}()

	go func() {
		err := b.client.StartConsumer(
			b.cfg.Service.Message.DocumentExcelDelete,
			b.messages[b.cfg.Service.Message.DocumentExcelDelete],
			true)
		if err != nil {
			errors <- err
		}
	}()

	go func() {
		err := b.client.StartConsumer(
			b.cfg.Service.Message.DocumentTXTDelete,
			b.messages[b.cfg.Service.Message.DocumentTXTDelete],
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
	err := b.client.CreateExchange(b.cfg.Service.Message.DocumentTXTSend)
	if err != nil {
		return err
	}
	err = b.client.CreateExchange(b.cfg.Service.Message.DocumentTXTGet)
	if err != nil {
		return err
	}
	err = b.client.CreateExchange(b.cfg.Service.Message.DocumentExcelGet)
	if err != nil {
		return err
	}
	err = b.client.CreateExchange(b.cfg.Service.Message.DocumentExcelSend)
	if err != nil {
		return err
	}
	err = b.client.CreateExchange(b.cfg.Service.Message.DocumentExcelDelete)
	if err != nil {
		return err
	}
	err = b.client.CreateExchange(b.cfg.Service.Message.DocumentTXTDelete)
	if err != nil {
		return err
	}
	return nil
}

func (b *Broker) CreateQueues() error {
	err := b.client.CreateQueue(b.cfg.Service.Message.DocumentExcelGet, true)
	if err != nil {
		return err
	}
	err = b.client.CreateQueue(b.cfg.Service.Message.DocumentTXTGet, true)
	if err != nil {
		return err
	}
	err = b.client.CreateQueue(b.cfg.Service.Message.DocumentExcelDelete, true)
	if err != nil {
		return err
	}
	err = b.client.CreateQueue(b.cfg.Service.Message.DocumentTXTDelete, true)
	if err != nil {
		return err
	}
	err = b.client.CreateQueue(b.cfg.Service.Message.DocumentExcelSend, true)
	if err != nil {
		return err
	}
	err = b.client.CreateQueue(b.cfg.Service.Message.DocumentTXTSend, true)
	if err != nil {
		return err
	}
	return nil
}
