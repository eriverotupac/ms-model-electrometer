package services

import (
	"context"
	"ms-model-electrometer/internal/models"
)

type IService interface {
	GetInfo(ctx context.Context, model string) (*models.ElectrometerResponse, error)
}
