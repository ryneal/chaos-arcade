package api

import (
	"html/template"
	"net/http"
	"path"
)

// Index godoc
// @Summary Index
// @Description renders chaos-arcade UI
// @Tags HTTP API
// @Produce html
// @Router / [get]
// @Success 200 {string} string "OK"
func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index.html").ParseFiles(path.Join(s.config.UIPath, "index.html"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(path.Join(s.config.UIPath, "index.html") + err.Error()))
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
		http.Error(w, path.Join(s.config.UIPath, "index.html")+err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) invadersHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("invaders.html").ParseFiles(path.Join(s.config.UIPath, "invaders.html"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(path.Join(s.config.UIPath, "invaders.html") + err.Error()))
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
		http.Error(w, path.Join(s.config.UIPath, "invaders.html")+err.Error(), http.StatusInternalServerError)
	}
}
