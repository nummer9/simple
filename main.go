package main

import (
	"github.com/BloodyRainer/simple/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
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
	http.Handle("/custom-variable", handlers.CustomVarHandler{})

	http.ListenAndServe(":" + strconv.Itoa(port), nil)

}
