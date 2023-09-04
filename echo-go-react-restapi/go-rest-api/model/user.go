package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email,max=30"`
	Password  string    `json:"password" validate:"required,min=6,max=30"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique"`
}
