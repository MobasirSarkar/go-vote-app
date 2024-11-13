package database

import (
	"database/sql"
	"fmt"

	"github.com/MobasirSarkar/go-vote-app/internal/models"
	"github.com/MobasirSarkar/go-vote-app/internal/utils"
)

// this is for users table queries
// AddUsers - it's add user data to users table
// FindUserByEmail - it's finds the user by email
// FindUserById - it's finds the user by id
// UpdataUsers  - update the user data in users table
// DeleteUsers  - shallow delete the user from the users table

func (s *service) AddUsers(u *models.User) error {
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashPassword
	query := `INSERT INTO users (name, email, password, role) VALUES($1, $2, $3, $4)`
	_, err = s.db.Exec(query, u.Name, u.Email, u.Password, u.Role)
	return err
}

func (s *service) FindUserByEmail(email string) (*models.User, error) {
	var u models.User
	query := `SELECT user_id, name, email, password, role FROM users WHERE email = $1`
	err := s.db.QueryRow(query, email).Scan(&u.UserId, &u.Name, &u.Email, &u.Password, &u.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with Email %s not found", email)
		}
		return nil, err
	}

	return &u, err
}
