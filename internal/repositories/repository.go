package repositories

import (
	"ms-model-electrometer/internal/models"

	"github.com/jmoiron/sqlx"
)

type IRepository interface {
	GetElectrometerInfo(dbConn *sqlx.DB, periodo string, sucursal string, zona string) ([]models.ElectrometerResponse, error)
	GetDatabaseConnectionString(sucursal string, codigoSistema string) (string, error)
}
