package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Person struct {
	Id      uuid.UUID `json:"id"`
	UserId  uint32    `json:"userId"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Email   string    `json:"email"`
}

func (person *Person) BeforeCreate(scope *gorm.DB) error {
	person.Id = uuid.New()
	return nil
}
