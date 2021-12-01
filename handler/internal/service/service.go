package service

import (
	"encoding/json"
	"github.com/arielkka/fallbox/handler/config"
	"github.com/arielkka/fallbox/handler/internal/models"
	"github.com/google/uuid"
)

type Service struct {
	cfg     *config.Config
	broker Broker
	storage Storage
}

func NewService(cfg *config.Config, broker Broker, storage Storage) *Service {
	return &Service{
		cfg:     cfg,
		broker: broker,
		storage: storage,
	}
}

func (s *Service) GetUser(login, password string) (string,error) {
	id, err := s.storage.GetUser(login, password)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *Service) CreateUser(login, password string) (string, error) {
	id := uuid.New()
	err := s.storage.CreateUser(login, password, id.String())
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (s *Service) GetUserPNG(userID, pngID string) ([]byte, error) {
	correlationID := uuid.New()
	request := &models.Request{
		UserID: userID,
		PngID:  pngID,
		JpgID:  "",
	}
	requestJSON, err := json.Marshal(request)
	if err	!= nil {
		return nil, err
	}
	err = s.broker.Publish(s.cfg.Service.Message.DocumentPNGGet, correlationID.String(), requestJSON)
	if err != nil {
		return nil, err
	}
	response, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentPNGGet, correlationID.String())

	png := new(models.PNG)
	err = json.Unmarshal(response, png)
	if err != nil {
		return nil, err
	}
	return png.Body, nil
}

func (s *Service) AddUserPNG(userID string, png *models.PNG) (string, error) {
	panic("implement me")
}

func (s *Service) DeleteUserPNG(userID, pngID string) error {
	panic("implement me")
}

func (s *Service) GetUserJPG(userID, jpgID string) ([]byte, error) {
	correlationID := uuid.New()
	request := &models.Request{
		UserID: userID,
		PngID:  "",
		JpgID:  jpgID,
	}
	requestJSON, err := json.Marshal(request)
	if err	!= nil {
		return nil, err
	}
	err = s.broker.Publish(s.cfg.Service.Message.DocumentJPGGet, correlationID.String(), requestJSON)
	if err != nil {
		return nil, err
	}
	response, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentJPGGet, correlationID.String())

	png := new(models.JPG)
	err = json.Unmarshal(response, png)
	if err != nil {
		return nil, err
	}
	return png.Body, nil
}

func (s *Service) AddUserJPG(userID string, jpg *models.JPG) (string, error) {
	panic("implement me")
}

func (s *Service) DeleteUserJPG(userID, pngID string) error {
	panic("implement me")
}