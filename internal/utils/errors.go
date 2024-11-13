package utils

import "errors"

var (
	ERROR_SAME_EMAIL          = errors.New("Email Already In Use")
	ERROR_NO_REPONSE          = errors.New("Response Is Empty")
	ERROR_INVALID_CREDENTIALS = errors.New("Invalid Credentials")
)
