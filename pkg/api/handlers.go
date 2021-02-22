package api

import (
	"html/template"
	"net/http"
	"path"
)

func (s *Server) templateHandler(file string, w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New(file).ParseFiles(path.Join(s.config.UIPath, file))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(path.Join(s.config.UIPath, file) + err.Error()))
		return
	}

	data := struct {
		Title string
		Logo  string
	}{
		Title: s.config.Hostname,
		Logo:  s.config.UILogo,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, path.Join(s.config.UIPath, file)+err.Error(), http.StatusInternalServerError)
	}
}

// Index godoc
// @Summary Index
// @Description renders chaos-arcade UI
// @Tags HTTP API
// @Produce html
// @Router / [get]
// @Success 200 {string} string "OK"
func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	s.templateHandler("index.html", w, r)
}

func (s *Server) snakeHandler(w http.ResponseWriter, r *http.Request) {
	s.templateHandler("snake.html", w, r)
}

func (s *Server) invadersHandler(w http.ResponseWriter, r *http.Request) {
	s.templateHandler("invaders.html", w, r)
}
