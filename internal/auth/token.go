package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/MobasirSarkar/go-vote-app/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

var secret = os.Getenv("JWT_SECRET")

func GenerateToken(user *models.User) (*LoginResponse, error) {

	// Create the Token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set Claims
	// Claims Store info for client to use for
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.UserId
	claims["name"] = user.Name
	claims["admin"] = CheckAdmin(user)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Returns a encoded token as a response
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = &user.UserId
	rtClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	rt, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token:        t,
		RefreshToken: rt,
	}, nil
}

// CheckAdmin check if the current is admin or not
func CheckAdmin(u *models.User) bool {
	if u.Role == "admin" {
		return true
	}
	return false
}

func ValidateToken(accessToken string) (models.User, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected siging method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	user := models.User{}
	if err != nil {
		return user, err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		user.UserId = payload["sub"].(string)

		return user, errors.New("Invalid Token")

	}

	return user, nil
}

func RefreshToken(refreshToken string) (*LoginResponse, error) {
	user, err := ValidateToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("Invalid refresh token: %v", err)
	}

	newTokens, err := GenerateToken(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate new token: %v", err)
	}
	return newTokens, nil
}
