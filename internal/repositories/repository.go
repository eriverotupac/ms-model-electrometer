package repositories

import (
	"context"
	"ms-model-electrometer/internal/models"

	"github.com/jmoiron/sqlx"
)

// IRepository definition
type IRepository interface {
	GetElectrometerInfo(ctx context.Context, dbConn *sqlx.DB, periodo string, sucursal string, zona string) ([]models.ElectrometerResponse, error)
	GetDatabaseConnectionString(ctx context.Context, sucursal string) (string, error)
}
