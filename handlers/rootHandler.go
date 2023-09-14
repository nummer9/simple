package handlers

import (
	"html/template"
	"net/http"
	"sync"

	"golang.org/x/exp/slog"
)

const indexTemplate = "index.gohtml"

var (
	tpl  *template.Template
	once sync.Once
)

type RootHandler struct{}

func (rcv RootHandler) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	once.Do(func() {
		tpl = template.Must(template.ParseFiles(indexTemplate))
	})

	slog.Info("received web-request", slog.String("route", "/"))

	data := struct {
		Title string
		H1    string
	}{
		Title: "simple",
		H1:    "42",
	}

	err := tpl.Execute(w, data)
	if err != nil {
		slog.Error("error writing to response writer", slog.String("route", "/"), slog.String("error", err.Error()))
	}
}
