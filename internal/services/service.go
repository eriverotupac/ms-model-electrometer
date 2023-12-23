package services

import (
	"context"
	"ms-model-electrometer/internal/models"
)

type IService interface {
	GetInfo(ctx context.Context, num string, sucursal string, zona string) ([]models.ElectrometerResponse, error)
}
