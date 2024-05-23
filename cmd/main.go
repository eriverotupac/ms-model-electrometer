package main

import (
	"fmt"
	"ms-model-electrometer/internal/api"
	"ms-model-electrometer/internal/config"
	"ms-model-electrometer/internal/repositories"
	"ms-model-electrometer/internal/server"
	"ms-model-electrometer/internal/services"
	"ms-model-electrometer/internal/utils"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	configs := config.GetVariables()

	cipher := utils.NewCipher(logger.Sugar(), *configs)

	httpServer := server.NewHTTPServer(logger.Sugar(), *configs)

	database := config.SetupDatabase(configs, logger.Sugar())
	err := database.Ping()
	if err != nil {
		fmt.Println("error in database connection")
	}
	defer database.Close()
	//Repos
	mainRepo := repositories.NewDatabaseRepository(logger.Sugar(), database)

	//Services
	mainSvc := services.NewDefaultService(logger.Sugar(), mainRepo, *cipher, *configs)

	//Controllers
	api.NewMainController(httpServer, mainSvc)

	httpServer.Start()
}
