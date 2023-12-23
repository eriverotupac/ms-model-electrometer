package repositories

import (
	"context"
	"ms-model-electrometer/internal/models"
)

// IRepository definition
type IRepository interface {
	GetElectrometerInfo(ctx context.Context, num string, sucursal string, zona string) ([]models.ElectrometerResponse, error)
}
