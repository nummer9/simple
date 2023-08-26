package handlers

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type HostnameHandler struct{}

func (rcv HostnameHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {

	log.WithFields(log.Fields{
		"route": "/hostname",
	}).Info("received web-request")

	hn := os.Getenv("HOSTNAME")

	if hn == "" {
		rw.Write([]byte("HOSTNAME is not set \n"))
	} else {
		rw.Write([]byte("HOSTNAME is " + hn + "\n"))
	}

}
