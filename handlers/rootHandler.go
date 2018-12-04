package handlers

import (
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"sync"
)

const indexTemplate = "handlers/index.gohtml"

var (
	tpl  *template.Template
	once sync.Once
)

type RootHandler struct {
}

func (rcv RootHandler) ServeHTTP(w http.ResponseWriter, rq *http.Request) {

	once.Do(func() {
		tpl = template.Must(template.ParseFiles(indexTemplate))

	})

	log.Info("received web-request to /")

	w.WriteHeader(http.StatusOK)

	data := struct {
		Title string
		H1    string
	}{
		Title: "simple",
		H1:    "42",
	}

	err := tpl.Execute(w, data)
	if err != nil {
		log.WithError(err).Error("error writing to response writer in root handler")
	}

}
