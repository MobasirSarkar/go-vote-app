package models

import "time"

type User struct {
	UserId    string    `json:"user_id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"email,required"`
	Role      string    `json:"role" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
