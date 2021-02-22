package api

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"math/rand"
	"net/http"
	"strings"
)

func (s *Server) randomPodHandler(w http.ResponseWriter, r *http.Request) {
	pod, err := s.getRandomPod()
	if err == nil {
		s.JSONResponse(w, r, pod)
	}
}

func (s *Server) randomPodDeleteHandler(w http.ResponseWriter, r *http.Request) {
	pod, err := s.getRandomPod()
	if err == nil {
		s.k8sClient.CoreV1().Pods(pod.Namespace).Delete(context.TODO(), pod.Name, metav1.DeleteOptions{})
		s.JSONResponse(w, r, pod)
	}
}

func contains(pod v1.Pod, namespaces []string) bool {
	for _, ns := range namespaces {
		if strings.Compare(ns, pod.Namespace) == 0 {
			return true
		}
	}
	return false
}

func extractValidPods(pods []v1.Pod, namespaces []string) []v1.Pod {
	var newPods []v1.Pod
	for _, p := range pods {
		if !contains(p, namespaces) {
			newPods = append(newPods, p)
		}
	}
	return newPods
}

func (s *Server) getRandomPod() (PodResponse, error) {
	excludedNamespaces := strings.Split(s.config.ExcludedNamespaces, ",")
	p, err := s.k8sClient.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	pods := extractValidPods(p.Items, excludedNamespaces)
	min := 0
	max := len(pods)
	i := rand.Intn(max-min) + min
	return PodResponse{
			Name:      pods[i].Name,
			Namespace: pods[i].Namespace,
		},
		nil
}

type PodResponse struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
