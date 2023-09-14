package main

import (
	"net/http"
	"strconv"

	"golang.org/x/exp/slog"

	"simple/handlers"
)

const port = 8080

func main() {
	slog.Info("Simple webserver started", slog.Int("port", port))

	http.Handle("/", handlers.RootHandler{})
	http.Handle("/health", handlers.HealthHandler{})
	http.Handle("/hostname", handlers.HostnameHandler{})
	http.Handle("/random-wait", handlers.RandomWaitHandler{})

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
