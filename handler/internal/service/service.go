package service

import "github.com/arielkka/fallbox/handler/config"

type Service struct {
	cfg     *config.Config
	storage Storage
}

func NewService(cfg *config.Config, storage Storage) *Service {
	return &Service{
		cfg:     cfg,
		storage: storage,
	}
}
