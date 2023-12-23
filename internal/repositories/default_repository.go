package repositories

import (
	"context"
	"database/sql"
	"errors"
	"ms-model-electrometer/internal/models"
	"strings"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type DatabaseRepository struct {
	log *zap.SugaredLogger
	db  *sqlx.DB
}

func NewDatabaseRepository(logger *zap.SugaredLogger, db *sqlx.DB) *DatabaseRepository {
	return &DatabaseRepository{
		log: logger,
		db:  db,
	}
}

func (r *DatabaseRepository) GetElectrometerInfo(ctx context.Context, elecNumber string, sucursal string, zona string) ([]models.ElectrometerResponse, error) {
	query := `EXEC sp_miStoreProceduree 'input_var', 'sucursal', 'zona'`
	query = strings.Replace(query, "input_var", elecNumber, -1)
	query = strings.Replace(query, "sucursal", sucursal, -1)
	query = strings.Replace(query, "zona", zona, -1)

	rows, err := r.db.Query(query)

	results := []models.ElectrometerResponse{}

	defer rows.Close()
	for rows.Next() {
		electroResponse := models.ElectrometerResponse{}
		err = rows.Scan(
			&electroResponse.CodigoSuministro,
			&electroResponse.CodigoMedidor,
			&electroResponse.NombreMarca,
			&electroResponse.NombreModelo,
			&electroResponse.DigitosDecimal,
			&electroResponse.LecturaActual)
		if err != nil {
			r.log.Error(err)
		}
		results = append(results, electroResponse)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			r.log.Info("No records found")
			return nil, errors.New("no records found")
		} else {
			r.log.Error("Error: " + err.Error())
			return nil, err
		}
	}

	return results, nil
}
