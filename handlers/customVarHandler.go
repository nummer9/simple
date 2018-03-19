package handlers

import (
	"fmt"
	"net/http"
	"os"
)

type CustomVarHandler struct{}

func (rcv CustomVarHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {

	fmt.Println("received web-request to /custom-variable")

	cv := os.Getenv("CUSTOM_VARIABLE")

	rw.WriteHeader(http.StatusOK)

	if cv == "" {
		rw.Write([]byte("CUSTOM_VARIABLE is not set \n"))
	} else {
		rw.Write([]byte("CUSTOM_VARIABLE is " + cv + "\n"))
	}

}
