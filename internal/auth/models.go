package auth

// Structure for Credentials
type LoginRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

// Structure for Response
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type Token struct {
	RefreshToken string `json:"refresh_token"`
}
