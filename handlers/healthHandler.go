package handlers

import (
	"net/http"

	"golang.org/x/exp/slog"
)

type HealthHandler struct{}

func (rcv HealthHandler) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	slog.Info("received web-request", slog.String("route", "health"))

	w.Write([]byte("simple health check: success \n"))
}
