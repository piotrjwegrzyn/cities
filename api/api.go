package api

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"text/template"
)

//go:embed html/index.html
var templateFS embed.FS

type PageData struct {
	Amount string
}

type Crawler interface {
	Cities(prefix string) ([]string, error)
}

type API struct {
	c Crawler
	t *template.Template
}

func New(c Crawler) (*API, error) {
	tmpl, err := template.ParseFS(templateFS, "html/index.html")
	if err != nil {
		return nil, err
	}

	return &API{
		c: c,
		t: tmpl,
	}, nil
}

func (a API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data PageData

	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/":
		w.WriteHeader(http.StatusOK)

	case r.Method == http.MethodPost && r.URL.Path == "/":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "could not parse form", http.StatusInternalServerError)

			return
		}

		prefix := r.FormValue("name_prefix")

		cities, err := a.c.Cities(prefix)
		if err != nil {
			http.Error(w, "failed to get cities", http.StatusInternalServerError)

			return
		}

		data.Amount = fmt.Sprint(len(cities))

	default:
		http.NotFound(w, r)

		return
	}

	if err := a.t.Execute(w, data); err != nil {
		slog.Error("error with rendering template", slog.String("error", err.Error()))
	}
}
