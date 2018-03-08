package handlers

import (
	"fmt"
	"net/http"
)

type RootHandler struct{}

func (rcv RootHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {

	fmt.Println("received web-request to /")

	rw.WriteHeader(http.StatusOK)

	rw.Write([]byte("This is a simple webservice"))

}
