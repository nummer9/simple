package main

import (
	"net/http"
	"fmt"
	"github.com/BloodyRainer/simple/handlers"
	"strconv"
)

const port = 8080

func main() {

	fmt.Println("Simple webserver is listening on port: " + strconv.Itoa(port))

	http.Handle("/health", handlers.HealthHandler{})
	http.Handle("/hostname", handlers.HostnameHandler{})
	http.ListenAndServe(":" + strconv.Itoa(port), nil)

}
