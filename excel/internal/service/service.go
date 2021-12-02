package service

import (
	"encoding/json"
	"sync"

	"github.com/arielkka/fallbox/excel/config"
	"github.com/arielkka/fallbox/excel/internal/models"
	logger "github.com/arielkka/fallbox/excel/pkg/logrus"
)

var mutex = sync.Mutex{}

type Service struct {
	cfg     *config.Config
	broker  Broker
	storage *Storage
	logger  logger.Logger
}

func NewService(cfg *config.Config, broker Broker, storage *Storage, log logger.Logger) *Service {
	return &Service{
		cfg:     cfg,
		broker:  broker,
		storage: storage,
		logger:  log,
	}
}

func (s *Service) Run() {
	stop := make(chan struct{})

	go func() {
		for {
			message, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentExcelSend)
			if err != nil {
				s.logger.Errorf("error = %s", err)
			}
			body, err := s.AddExcel(message.Body)
			if err != nil {
				s.logger.Errorf("error = %s", err)
				err = s.broker.Respond(message.ReplyTo, s.cfg.Service.Message.DocumentExcelSend, message.ID, body)
				if err != nil {
					s.logger.Fatalf("error = %s", err)
				}
			} else {
				err = s.broker.Respond(message.ReplyTo, s.cfg.Service.Message.DocumentExcelSend, message.ID, body)
				if err != nil {
					s.logger.Fatalf("error = %s", err)
				}
			}
		}
	}()

	go func() {
		for {
			message, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentExcelGet)
			if err != nil {
				s.logger.Errorf("error = %s", err)
			}
			body, err := s.GetExcel(message.Body)
			if err != nil {
				s.logger.Errorf("error = %s", err)
				err = s.broker.Respond(message.ReplyTo, s.cfg.Service.Message.DocumentExcelGet, message.ID, body)
				if err != nil {
					s.logger.Fatalf("error = %s", err)
				}
			} else {
				err = s.broker.Respond(message.ReplyTo, s.cfg.Service.Message.DocumentExcelGet, message.ID, body)
				if err != nil {
					s.logger.Fatalf("error = %s", err)
				}
			}
		}
	}()

	go func() {
		for {
			message, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentExcelDelete)
			if err != nil {
				s.logger.Errorf("error = %s", err)
			}
			body, err := s.DeleteExcel(message.Body)
			if err != nil {
				s.logger.Errorf("error = %s", err)
				err = s.broker.Respond(message.ReplyTo, s.cfg.Service.Message.DocumentExcelDelete, message.ID, body)
				if err != nil {
					s.logger.Fatalf("error = %s", err)
				}
			} else {
				err = s.broker.Respond(message.ReplyTo, s.cfg.Service.Message.DocumentExcelDelete, message.ID, body)
				if err != nil {
					s.logger.Fatalf("error = %s", err)
				}
			}
		}
	}()
	<-stop
}

func (s *Service) AddExcel(body []byte) ([]byte, error) {
	var req models.Request
	err := json.Unmarshal(body, &req)
	if err != nil {
		return nil, err
	}
	id, err := s.storage.AddExcel(req.UserID, req.Body)
	var res models.Response
	res.ID = id
	result, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) GetExcel(body []byte) ([]byte, error) {
	var req models.Request
	err := json.Unmarshal(body, &req)
	if err != nil {
		return nil, err
	}
	excel, err := s.storage.GetExcel(req.UserID, req.ID)
	var res models.Response
	res.Body = excel
	result, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) DeleteExcel(body []byte) ([]byte, error) {
	var req models.Request
	err := json.Unmarshal(body, &req)
	if err != nil {
		return nil, err
	}
	err = s.storage.DeleteExcel(req.UserID, req.ID)
	var res models.IsDeleted
	if err != nil {
		res.Flag = false
		result, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	res.Flag = true
	result, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return result, nil
}
