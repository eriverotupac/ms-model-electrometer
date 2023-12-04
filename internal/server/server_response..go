package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
)

func RenderJSON(ctx context.Context, w http.ResponseWriter, httpStatusCode int, payload interface{}) {
	w.Header().Set(middleware.RequestIDHeader, middleware.GetReqID(ctx))
	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(httpStatusCode)
	_, _ = w.Write(js)
}

func RenderError(ctx context.Context, w http.ResponseWriter, err error) {
	var payload = map[string]string{}
	if strings.Contains(err.Error(), "records found") {
		payload = map[string]string{
			"code":    strconv.Itoa(http.StatusNotFound),
			"message": err.Error(),
		}
	} else {
		payload = map[string]string{
			"code":    strconv.Itoa(http.StatusInternalServerError),
			"message": "something went wrong...." + err.Error(),
		}
	}
	RenderJSON(ctx, w, http.StatusInternalServerError, payload)
}
