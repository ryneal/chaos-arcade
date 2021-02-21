package api

import (
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"math/rand"
	"net/http"
)

func (s *Server) randomPodHandler(w http.ResponseWriter, r *http.Request) {
	s.JSONResponse(w, r, s.getRandomPod())
}

func (s *Server) randomPodDeleteHandler(w http.ResponseWriter, r *http.Request) {
	pod := s.getRandomPod()
	s.k8sClient.CoreV1().Pods(pod.Namespace).Delete(context.TODO(), pod.Name, v1.DeleteOptions{})
	w.WriteHeader(http.StatusAccepted)
}

func (s *Server) getRandomPod() PodResponse {
	p, err := s.k8sClient.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	pods := p.Items
	min := 0
	max := len(pods)
	i := rand.Intn(max-min) + min
	return PodResponse{
		Name:      pods[i].Name,
		Namespace: pods[i].Namespace,
	}
}

type PodResponse struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
