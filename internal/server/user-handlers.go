package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MobasirSarkar/go-vote-app/internal/models"
	"github.com/MobasirSarkar/go-vote-app/internal/utils"
)

func (s *Server) HandlerAddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var details utils.DuplicateEntryError

	// Decodes the reponse in user
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error Decoding response body: %s", err)
		http.Error(w, "Please provide the correct input", http.StatusBadRequest)
		return
	}

	// validate the user response
	utils.InitValidator()
	if err := utils.Validate(&user); err != nil {
		log.Printf("Format Error: %s", err)
		http.Error(w, "Validation Error. Please Check Your Input", http.StatusBadRequest)
		return
	}

	// add reponse to the table users
	err := s.db.AddUsers(&user)
	if err != nil {
		if utils.IsDuplicateEntryError(err, &details) {
			log.Printf("Unique Key Violation: %s", err)
			http.Error(w, details.Error(), http.StatusConflict)
			return
		} else {
			log.Printf("Server Error %s", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}
	}

	// Status Ok if succeed
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User Created Successfully")

}
