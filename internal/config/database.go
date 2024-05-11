package config

import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const DATABASE_URL_EXAMPLE = "sqlserver://<usuario>:<password>@<dominio>:<puerto>?database=master&connection+timeout=30"

func SetupDatabase(env *Environment, logger *zap.SugaredLogger) *sqlx.DB {
	logger.Info("start connection to database")
	fmt.Println(env.DatabaseUrl)

	dbConn, err := sqlx.Connect("sqlserver", env.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	return dbConn
}
