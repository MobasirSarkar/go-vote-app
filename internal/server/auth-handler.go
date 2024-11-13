package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MobasirSarkar/go-vote-app/internal/auth"
	"github.com/MobasirSarkar/go-vote-app/internal/utils"
)

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var loginRequest auth.LoginRequest
	// Decodes the reponse in user
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Printf("Error Decoding response body: %s", err)
		http.Error(w, "Please provide the correct input", http.StatusBadRequest)
		return
	}

	user, err := s.db.FindUserByEmail(loginRequest.Email)
	if err != nil {
		log.Print(err)
		http.Error(w, "Invalid Credentials", http.StatusConflict)
		return
	}

	if !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
		log.Print(err)
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	tokens, err := auth.GenerateToken(user)
	if err != nil {
		log.Print(err)
		http.Error(w, "Unable to Generate Token", http.StatusUnauthorized)
		return
	}

	var token = &auth.LoginResponse{
		Token:        tokens.Token,
		RefreshToken: tokens.RefreshToken,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
