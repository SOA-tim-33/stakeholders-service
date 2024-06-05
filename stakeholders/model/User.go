package model

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     UserRole  `json:"role"`
	IsActive bool      `json:"isActive"`
	Email    string    `json:"email"`
	Token    string    `json:"token"`
}

type UserRole string

const (
	Administrator UserRole = "administrator"
	Author        UserRole = "author"
	Tourist       UserRole = "tourist"
)

func (user *User) GetPrimaryRoleName() string {
	return strings.ToLower(string(user.Role))
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.Id = uuid.New()
	return nil
}
