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

func (r *DatabaseRepository) GetElectrometerInfo(ctx context.Context, elecNumber string) (*models.ElectrometerResponse, error) {
	query := `EXEC sp_miStoreProcedure 'input_var'`
	query = strings.Replace(query, "input_var", elecNumber, -1)

	var codigoSuministro string
	var codigoMedidor string
	var nombreMarca string
	var nombreModelo string
	var digitosDecimal string
	var lecturaActual string

	row := r.db.QueryRowContext(ctx, query, elecNumber)

	err := row.Scan(
		&codigoSuministro,
		&codigoMedidor,
		&nombreMarca,
		&nombreModelo,
		&digitosDecimal,
		&lecturaActual)

	if err != nil {
		if err == sql.ErrNoRows {
			r.log.Info("No records found")
			return nil, errors.New("no records found")
		} else {
			r.log.Error("Error: " + err.Error())
			return nil, err
		}
	}

	return &models.ElectrometerResponse{
		CodigoSuministro: codigoSuministro,
		CodigoMedidor:    codigoMedidor,
		NombreMarca:      nombreMarca,
		NombreModelo:     nombreModelo,
		LecturaActual:    lecturaActual,
	}, nil
}
