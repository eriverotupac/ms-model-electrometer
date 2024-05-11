package services

import (
	"context"
	"ms-model-electrometer/internal/models"
	"ms-model-electrometer/internal/repositories"

	"github.com/jmoiron/sqlx"
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

func (s *DefaultService) GetInfo(ctx context.Context, periodo string, sucursal string, zona string) ([]models.ElectrometerResponse, error) {
	databaseUrl, err := s.electrometerRepo.GetDatabaseConnectionString(ctx, sucursal)
	if err != nil {
		s.log.Errorf("failed to get the connection string: %v", err.Error())
		return nil, err
	}

	dbConnection, err := sqlx.Connect("sqlserver", databaseUrl)

	if err != nil {
		s.log.Errorf("failed to connect: %v", err.Error())
	}
	defer dbConnection.Close()

	electrometer, err := s.electrometerRepo.GetElectrometerInfo(ctx, dbConnection, periodo, sucursal, zona)
	if err != nil {
		s.log.Errorf("failed to get the electrometer data: %v", err.Error())
		return nil, err
	}
	return electrometer, nil
}
