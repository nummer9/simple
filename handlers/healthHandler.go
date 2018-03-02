package handlers

import (
	"net/http"
	"fmt"
)

type HealthHandler struct{}

func (rcv HealthHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {

	fmt.Println("received web-request to /health")

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("simple health check: success \n"))
}
