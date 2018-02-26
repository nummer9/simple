package simple

import (
	"net/http"
	"fmt"
)

type SimpleHealthHandler struct{}

func (rcv SimpleHealthHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {

	fmt.Println("received web-request to /health")

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("simple health check\n"))
}
