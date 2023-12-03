package server

import (
	"ms-model-electrometer/internal/config"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	Logger *zap.SugaredLogger
	sc     config.Environment
	Router *chi.Mux
}

func NewHTTPServer(logger *zap.SugaredLogger, serverConf config.Environment) *HTTPServer {
	router := chi.NewRouter()

	router.Use(middleware.Timeout(60 * time.Second))
	return &HTTPServer{
		Logger: logger,
		sc:     serverConf,
		Router: chi.NewRouter(),
	}
}

func (srv *HTTPServer) Start() {
	listeningAddr := ":" + srv.sc.AppPort
	srv.Logger.Infof("Server listening on port %s", listeningAddr)

	err := http.ListenAndServe(listeningAddr, srv.Router)
	if err != nil {
		srv.Logger.Fatalf("Failed to start http server. %v", err)
	}
}
