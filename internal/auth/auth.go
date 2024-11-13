package auth

import (
	"net/http"

	"github.com/MobasirSarkar/go-vote-app/internal/models"
)

type Auth struct{}

func (auth *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var LoginRequest LoginRequest
	var user *models.User

   user, err :=


}

