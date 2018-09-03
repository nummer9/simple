package handlers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type CustomVarHandler struct{}

func (rcv CustomVarHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {

	log.WithFields(log.Fields{
		"route" : "/custom-variable",
	}).Info("received web-request")

	cv := os.Getenv("CUSTOM_VARIABLE")

	rw.WriteHeader(http.StatusOK)

	if cv == "" {
		rw.Write([]byte("CUSTOM_VARIABLE is not set \n"))
	} else {
		rw.Write([]byte("CUSTOM_VARIABLE is " + cv + "\n"))
	}

}
