package api

import (
	"encoding/json"
	"ms-model-electrometer/internal/server"
	"ms-model-electrometer/internal/services"
	"net/http"

	"go.uber.org/zap"
)

type MainController struct {
	mainSvc services.IService
	log     *zap.SugaredLogger
}

func NewMainController(srv *server.HTTPServer, ms services.IService) *MainController {
	mc := &MainController{
		log:     srv.Logger,
		mainSvc: ms,
	}

	srv.Router.Get("/electrometer_model", mc.getElectrometerInfo)

	return mc
}

func (mc *MainController) getElectrometerInfo(w http.ResponseWriter, r *http.Request) {
	mc.log.Info("Incoming request to get electrometer info")

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var electrometerNumber string
	err := dec.Decode(&electrometerNumber)

	elecRecord, err := mc.mainSvc.GetInfo(r.Context(), electrometerNumber)

	if err != nil {

		server.RenderError(r.Context(), w, err)
		return
	}
	server.RenderJSON(r.Context(), w, http.StatusOK, elecRecord)
}
