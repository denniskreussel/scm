package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	fs := http.FileServer(http.Dir("web/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", s.indexHandler)

	return r
}

type IndexData struct {
	Content string
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {

	data := IndexData{
		Content: "Hello World",
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	err = tmpl.ExecuteTemplate(w, "Content", data)
	if err != nil {
		log.Fatalf("Could not execute template. Err: %v", err)
	}
	if err != nil {
		log.Fatalf("Could not execute template. Err: %v", err)
	}
}
