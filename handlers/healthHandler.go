package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HealthHandler struct{}

func (rcv HealthHandler) ServeHTTP(w http.ResponseWriter, rq *http.Request) {

	log.WithFields(log.Fields{
		"route": "/health",
	}).Info("received web-request")

	w.Write([]byte("simple health check: success \n"))
}
