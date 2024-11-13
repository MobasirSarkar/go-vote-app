package utils

import (
	"github.com/MobasirSarkar/go-vote-app/internal/models"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

func Validate(u *models.User) error {
	return validate.Struct(u)
}
