package service

import (
	"encoding/json"

	"github.com/arielkka/fallbox/handler/config"
	"github.com/arielkka/fallbox/handler/internal/models"
	"github.com/arielkka/fallbox/handler/pkg/errors"

	"github.com/google/uuid"
)

type Service struct {
	cfg     *config.Config
	broker  Broker
	storage IStorage
}

func NewService(cfg *config.Config, broker Broker, storage IStorage) *Service {
	return &Service{
		cfg:     cfg,
		broker:  broker,
		storage: storage,
	}
}

func (s *Service) GetUser(login, password string) (string, error) {
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
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
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
	correlationID := uuid.New().String()

	request := &models.Request{
		Body: png.Body,
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	err = s.broker.Publish(s.cfg.Service.Message.DocumentPNGSend, correlationID, requestJSON)
	if err != nil {
		return "", err
	}

	response, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentPNGSend, correlationID)
	if err != nil {
		return "", err
	}

	pngID := new(models.PngID)
	err = json.Unmarshal(response, pngID)
	if err != nil {
		return "", err
	}
	return pngID.ID, nil
}

func (s *Service) DeleteUserPNG(userID, pngID string) error {
	correlationID := uuid.New().String()

	request := &models.Request{
		UserID: userID,
		PngID:  pngID,
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}

	err = s.broker.Publish(s.cfg.Service.Message.DocumentPNGDelete, correlationID, requestJSON)
	if err != nil {
		return err
	}

	response, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentPNGDelete, correlationID)
	if err != nil {
		return err
	}

	isDeleted := new(models.IsDeleted)
	err = json.Unmarshal(response, isDeleted)
	if err != nil {
		return err
	}
	if !isDeleted.Flag {
		return errors.NotFound()
	}
	return nil
}

func (s *Service) GetUserJPG(userID, jpgID string) ([]byte, error) {
	correlationID := uuid.New()
	request := &models.Request{
		UserID: userID,
		JpgID:  jpgID,
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
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
	correlationID := uuid.New().String()

	request := &models.Request{
		Body: jpg.Body,
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	err = s.broker.Publish(s.cfg.Service.Message.DocumentJPGSend, correlationID, requestJSON)
	if err != nil {
		return "", err
	}

	response, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentJPGSend, correlationID)
	if err != nil {
		return "", err
	}

	jpgID := new(models.JpgID)
	err = json.Unmarshal(response, jpgID)
	if err != nil {
		return "", err
	}
	return jpgID.ID, nil
}

func (s *Service) DeleteUserJPG(userID, jpgID string) error {
	correlationID := uuid.New().String()

	request := &models.Request{
		UserID: userID,
		JpgID:  jpgID,
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}

	err = s.broker.Publish(s.cfg.Service.Message.DocumentJPGDelete, correlationID, requestJSON)
	if err != nil {
		return err
	}

	response, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentJPGDelete, correlationID)
	if err != nil {
		return err
	}

	isDeleted := new(models.IsDeleted)
	err = json.Unmarshal(response, isDeleted)
	if err != nil {
		return err
	}
	if !isDeleted.Flag {
		return errors.NotFound()
	}
	return nil
}
