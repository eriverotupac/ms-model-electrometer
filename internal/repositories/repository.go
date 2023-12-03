package repositories

import (
	"context"
	"ms-model-electrometer/internal/models"
)

// IRepository definition
type IRepository interface {
	GetElectrometerInfo(ctx context.Context, model string) (*models.ElectrometerResponse, error)
}
