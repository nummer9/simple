package handlers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type HealthHandler struct{}

func (rcv HealthHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {

	log.WithFields(log.Fields{
		"route" : "/health",
	}).Info("received web-request")

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("simple health check: success \n"))
}
