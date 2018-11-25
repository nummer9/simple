package main

import (
	"net/http"
	"os"
	"simple/handlers"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

const port = 8080

func main() {

	log.SetOutput(os.Stdout)

	log.WithFields(log.Fields{
		"time": time.Now().Format(time.RFC3339),
		"port": strconv.Itoa(port),
	}).Info("Simple webserver started")

	http.Handle("/", handlers.RootHandler{})
	http.Handle("/health", handlers.HealthHandler{})
	http.Handle("/hostname", handlers.HostnameHandler{})

	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}
