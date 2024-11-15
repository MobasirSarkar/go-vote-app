package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /", s.HandlerHome)
	router.Handle("GET /ping", s.HandleAuth(http.HandlerFunc(s.HandlerPing)))

	router.HandleFunc("POST /create-user", s.HandlerAddUser)
	router.HandleFunc("POST /sign-in", s.LoginHandler)

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

func (s *Server) HandlerPing(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)

	err := s.db.Ping()
	if err != nil {
		resp["message"] = `Db error`
	} else {
		resp["message"] = "Database Connected"
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
