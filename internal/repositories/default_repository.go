package repositories

import (
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

func (r *DatabaseRepository) GetElectrometerInfo(dbConnection *sqlx.DB, periodo string, sucursal string, zona string) ([]models.ElectrometerResponse, error) {
	query := `EXEC sp_miStoreProcedure01 'input_01', 'input_02', 'input_03'`
	query = strings.Replace(query, "input_01", periodo, -1)
	query = strings.Replace(query, "input_02", sucursal, -1)
	query = strings.Replace(query, "input_03", zona, -1)

	rows, err := dbConnection.Query(query)

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

func (r *DatabaseRepository) GetDatabaseConnectionString(sucursal string, codigoSistema string) (string, error) {
	query := `EXEC _spObtenerContextoSeguridad 'cod_sistema', 'sucursal'`
	query = strings.Replace(query, "cod_sistema", codigoSistema, -1)
	query = strings.Replace(query, "sucursal", sucursal, -1)

	rows, err := r.db.Query(query)
	if err != nil {
		r.log.Error("Error: " + err.Error())
		return "", err
	}
	var ConexionBase string

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&ConexionBase)
		if err != nil {
			r.log.Error(err)
		}

	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Info("No records found")
			return "", errors.New("no records found")
		} else {
			r.log.Error("Error: " + err.Error())
			return "", err
		}
	}

	return ConexionBase, nil
}
