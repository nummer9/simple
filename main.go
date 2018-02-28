package main

import (
	"net/http"
	"fmt"
	"github.com/BloodyRainer/simple/handlers"
)

func main() {

	fmt.Println("starting simple webserver")

	http.Handle("/health", handlers.HealthHandler{})
	http.ListenAndServe(":8080", nil)

}
