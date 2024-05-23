package api

import (
	"ms-model-electrometer/internal/server"
	"ms-model-electrometer/internal/services"
	"net/http"

	"go.uber.org/zap"
)

type MainController struct {
	electrometerSvc services.IService
	log             *zap.SugaredLogger
}

func NewMainController(srv *server.HTTPServer, ms services.IService) *MainController {
	mc := &MainController{
		log:             srv.Logger,
		electrometerSvc: ms,
	}

	srv.Router.Get("/electro-model", mc.getElectrometerInfo)

	return mc
}

func (mc *MainController) getElectrometerInfo(w http.ResponseWriter, r *http.Request) {
	mc.log.Info("Incoming request to get electrometer info")

	query := r.URL.Query()
	periodo := query.Get("periodo")
	sucursal := query.Get("sucursal")
	zona := query.Get("zona")

	elecRecord, err := mc.electrometerSvc.GetInfo(periodo, sucursal, zona)

	if err != nil {
		server.RenderError(r.Context(), w, err)
		return
	}
	server.RenderJSON(r.Context(), w, http.StatusOK, elecRecord)
}
