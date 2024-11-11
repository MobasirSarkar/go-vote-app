package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", s.HandlerHome)

	return router
}

func (s *Server) HandlerHome(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)

	resp["message"] = "Hello Mobasir"

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
