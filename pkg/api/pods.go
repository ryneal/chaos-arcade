package api

import (
	"net/http"
)

func (s *Server) podsHandler(w http.ResponseWriter, r *http.Request) {
	data := [1]PodResponse{
		{
			Name:      "test",
			Namespace: "test-namespace",
		},
	}

	s.JSONResponse(w, r, data)
}

func (s *Server) randomPodHandler(w http.ResponseWriter, r *http.Request) {
	data := PodResponse{
		Name:      "test",
		Namespace: "test-namespace",
	}

	s.JSONResponse(w, r, data)
}

type PodResponse struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
