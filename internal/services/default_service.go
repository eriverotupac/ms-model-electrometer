package services

import (
	"fmt"
	"ms-model-electrometer/internal/config"
	"ms-model-electrometer/internal/models"
	"ms-model-electrometer/internal/repositories"
	"ms-model-electrometer/internal/utils"
	"strings"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const DATABASE_URL_EXAMPLE = "sqlserver://%s:%s@%s?database=%s&connection+timeout=30"

type DefaultService struct {
	log              *zap.SugaredLogger
	electrometerRepo repositories.IRepository
	cipher           utils.Cipher
	configs          config.Environment
}

func NewDefaultService(logger *zap.SugaredLogger, r repositories.IRepository, cipher utils.Cipher, env config.Environment) *DefaultService {
	return &DefaultService{
		log:              logger,
		electrometerRepo: r,
		cipher:           cipher,
		configs:          env,
	}
}

func (s *DefaultService) GetInfo(periodo string, sucursal string, zona string) ([]models.ElectrometerResponse, error) {
	codigoSistema := s.configs.SystemCode
	databaseUrlCiphered, err := s.electrometerRepo.GetDatabaseConnectionString(sucursal, codigoSistema)

	if err != nil {
		s.log.Errorf("failed to get the connection string: %s", err.Error())
		return nil, err
	}

	s.log.Infof("get connection string encrypted from db: %s", databaseUrlCiphered)

	databaseData, err := s.cipher.DecryptString(databaseUrlCiphered)
	if err != nil {
		s.log.Errorf("failed to decrypt database url: %v", err.Error())
		return nil, err
	}
	s.log.Infof("value got after decipher: %v", databaseData)

	dataBaseServerValues := strings.Split(databaseData, "|")

	if len(dataBaseServerValues) == 0 {
		s.log.Error("failed to parse database connection")
		return nil, err
	}

	connServer := dataBaseServerValues[0]
	ipAndServer := strings.Split(connServer, "\\")

	if len(ipAndServer) == 0 {
		s.log.Errorf("failed to parse database server from string connection")
		return nil, err
	}
	connIP := ipAndServer[0]
	connDatabaseName := dataBaseServerValues[1]

	databaseUrl := fmt.Sprintf(DATABASE_URL_EXAMPLE, s.configs.UserDB, s.configs.PasswordDB, connIP, connDatabaseName)

	s.log.Infof("connection string builded: %s", databaseUrl)

	dbConnection, err := sqlx.Connect("sqlserver", databaseUrl)

	if err != nil {
		s.log.Errorf("failed to connect: %v", err.Error())
		return nil, err
	}
	defer dbConnection.Close()

	electrometer, err := s.electrometerRepo.GetElectrometerInfo(dbConnection, periodo, sucursal, zona)
	if err != nil {
		s.log.Errorf("failed to get the electrometer data: %v", err.Error())
		return nil, err
	}
	return electrometer, nil
}
