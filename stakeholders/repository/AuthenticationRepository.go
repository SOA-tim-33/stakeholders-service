package repository

import (
	"gorm.io/gorm"
)

type AuthenticationRepository struct {
	DatabaseConnection *gorm.DB
}
