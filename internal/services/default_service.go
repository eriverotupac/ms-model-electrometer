package services

import (
	"context"
	"ms-model-electrometer/internal/models"
	"ms-model-electrometer/internal/repositories"

	"go.uber.org/zap"
)

type DefaultService struct {
	log              *zap.SugaredLogger
	electrometerRepo repositories.IRepository
}

func NewDefaultService(logger *zap.SugaredLogger, r repositories.IRepository) *DefaultService {
	return &DefaultService{
		log:              logger,
		electrometerRepo: r,
	}
}

func (s *DefaultService) GetInfo(ctx context.Context, model string) (*models.ElectrometerResponse, error) {

	electrometer, err := s.electrometerRepo.GetElectrometerInfo(ctx, model)
	if err != nil {
		s.log.Errorf("failed to get the electrometer data: %v", err.Error())
		return nil, err
	}
	return electrometer, nil
}
