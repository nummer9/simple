package handlers

import (
	"fmt"
	"net/http"
	"os"
)

type HostnameHandler struct{}

func (rcv HostnameHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {

	fmt.Println("received web-request to /hostname")

	hn := os.Getenv("HOSTNAME")

	rw.WriteHeader(http.StatusOK)

	if hn == "" {
		rw.Write([]byte("HOSTNAME is not set \n"))
	} else {
		rw.Write([]byte("HOSTNAME is " + hn + "\n"))
	}

}
