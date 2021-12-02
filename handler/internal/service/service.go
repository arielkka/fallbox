package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/arielkka/fallbox/handler/config"
	"github.com/arielkka/fallbox/handler/internal/models"
	myerrors "github.com/arielkka/fallbox/handler/pkg/errors"
	"github.com/google/uuid"
)

type Service struct {
	cfg     *config.Config
	broker  Broker
	storage *Storage
}

func NewService(cfg *config.Config, broker Broker, storage *Storage) *Service {
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

func (s *Service) GetUserTxt(userID string, txtID int) error {
	correlationID := uuid.New().String()
	request := &models.Request{
		UserID: userID,
		ID:     txtID,
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}
	err = s.broker.Publish(s.cfg.Service.Message.DocumentTXTGet, correlationID, requestJSON)
	if err != nil {
		return err
	}
	response, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentTXTGet, correlationID)

	txt := new(models.Response)
	err = json.Unmarshal(response, txt)
	if err != nil {
		return err
	}
	err = s.downloadTxt(txt)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AddUserTxt(userID string, text []byte) (int, error) {
	correlationID := uuid.New().String()

	request := &models.Request{
		UserID: userID,
		Body:   text,
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return -1, err
	}

	err = s.broker.Publish(s.cfg.Service.Message.DocumentTXTSend, correlationID, requestJSON)
	if err != nil {
		return -1, err
	}

	response, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentTXTSend, correlationID)
	if err != nil {
		return -1, err
	}

	txtID := new(models.Response)
	err = json.Unmarshal(response, txtID)
	if err != nil {
		return -1, err
	}
	return txtID.ID, nil
}

func (s *Service) DeleteUserTxt(userID string, txtID int) error {
	correlationID := uuid.New().String()

	request := &models.Request{
		UserID: userID,
		ID:     txtID,
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}

	err = s.broker.Publish(s.cfg.Service.Message.DocumentTXTDelete, correlationID, requestJSON)
	if err != nil {
		return err
	}

	response, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentTXTDelete, correlationID)
	if err != nil {
		return err
	}

	isDeleted := new(models.IsDeleted)
	err = json.Unmarshal(response, isDeleted)
	if err != nil {
		return err
	}
	if !isDeleted.Flag {
		return myerrors.NotFound()
	}
	return nil
}

func (s *Service) GetUserExcel(userID string, excelID int) error {
	correlationID := uuid.New().String()
	request := &models.Request{
		UserID: userID,
		ID:     excelID,
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}
	err = s.broker.Publish(s.cfg.Service.Message.DocumentExcelGet, correlationID, requestJSON)
	if err != nil {
		return err
	}
	response, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentExcelGet, correlationID)

	excel := new(models.Response)
	err = json.Unmarshal(response, excel)
	if err != nil {
		return err
	}
	err = s.downloadExcel(excel)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AddUserExcel(userID string, excel []byte) (int, error) {
	correlationID := uuid.New().String()

	request := &models.Request{
		UserID: userID,
		Body:   excel,
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return -1, err
	}

	err = s.broker.Publish(s.cfg.Service.Message.DocumentExcelSend, correlationID, requestJSON)
	if err != nil {
		return -1, err
	}

	response, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentExcelSend, correlationID)
	if err != nil {
		return -1, err
	}

	excelID := new(models.Response)
	err = json.Unmarshal(response, excelID)
	if err != nil {
		return -1, err
	}
	return excelID.ID, nil
}

func (s *Service) DeleteUserExcel(userID string, excelID int) error {
	correlationID := uuid.New().String()

	request := &models.Request{
		UserID: userID,
		ID:     excelID,
	}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}

	err = s.broker.Publish(s.cfg.Service.Message.DocumentExcelDelete, correlationID, requestJSON)
	if err != nil {
		return err
	}

	response, err := s.broker.Subscribe(s.cfg.Service.Message.DocumentExcelDelete, correlationID)
	if err != nil {
		return err
	}

	isDeleted := new(models.IsDeleted)
	err = json.Unmarshal(response, isDeleted)
	if err != nil {
		return err
	}
	if !isDeleted.Flag {
		return myerrors.NotFound()
	}
	return nil
}

func (s *Service) downloadExcel(resp *models.Response) error {
	buffer := bytes.NewBuffer(resp.Body)
	all, err := io.ReadAll(buffer)
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("../../output/%v.xlsx", resp.ID))
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	_, err = file.Write(all)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) downloadTxt(resp *models.Response) error {
	buffer := bytes.NewBuffer(resp.Body)
	all, err := io.ReadAll(buffer)
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("../../output/%v.txt", resp.ID))
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	_, err = file.Write(all)
	if err != nil {
		return err
	}
	return nil
}
