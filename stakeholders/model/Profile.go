package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	Id                     uuid.UUID `json:"id"`
	FirstName              string    `json:"firstName"`
	LastName               string    `json:"lastName"`
	ProfilePicture         string    `json:"profilePicture"`
	Biography              string    `json:"biography"`
	Motto                  string    `json:"motto"`
	UserId                 uuid.UUID `json:"userId"`
	IsActive               bool      `json:"isActive"`
	XP                     uint32    `json:"xP"`
	IsFirstPurchased       bool      `json:"isFirstPurchased"`
	QuestionnaireDone      bool      `json:"questionnaireDone"`
	NumberOfCompletedTours uint32    `json:"numberOfCompletedTours"`
	RequestSent            bool      `json:"requestSent"`
}

func (profile *Profile) BeforeCreate(scope *gorm.DB) error {
	profile.Id = uuid.New()
	return nil
}
