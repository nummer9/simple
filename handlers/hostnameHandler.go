package handlers

import (
	"net/http"
	"os"

	"golang.org/x/exp/slog"
)

type HostnameHandler struct{}

func (rcv HostnameHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	slog.Info("received web-request", slog.String("route", "/hostname"))

	hn := os.Getenv("HOSTNAME")

	if hn == "" {
		rw.Write([]byte("HOSTNAME is not set \n"))
	} else {
		rw.Write([]byte("HOSTNAME is " + hn + "\n"))
	}
}
