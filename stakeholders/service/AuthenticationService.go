package service

import (
	"errors"
	"stakeholders-service/dto"
	"stakeholders-service/model"
	"stakeholders-service/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthenticationService struct {
	UserRepository    *repository.UserRepository
	ProfileRepository *repository.ProfileRepository
	TokenGenerator    *repository.JwtGenerator
}

func NewAuthenticationService(userRepository *repository.UserRepository, tokenGenerator *repository.JwtGenerator, profileRepository *repository.ProfileRepository) *AuthenticationService {
	return &AuthenticationService{
		UserRepository:    userRepository,
		TokenGenerator:    tokenGenerator,
		ProfileRepository: profileRepository,
	}
}

func (service *AuthenticationService) RegisterTourist(account *dto.AccountRegistrationDto, token string) (*dto.AuthenticationTokensDto, error) {
	exists, err := service.UserRepository.Exists(account.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("non-unique username")
	}

	user := &model.User{
		Username: account.Username,
		Password: account.Pasysword,
		Role:     model.Tourist,
		Email:    account.Email,
		Token:    token,
		IsActive: true,
	}
	newUser, err := service.UserRepository.Save(user)
	if err != nil {
		return nil, err
	}

	profile := &model.Profile{
		FirstName:              account.Name,
		LastName:               account.Surname,
		ProfilePicture:         account.ProfilePicture,
		Biography:              account.Biography,
		Motto:                  account.Motto,
		UserId:                 newUser.Id,
		IsActive:               true,
		QuestionnaireDone:      false,
		NumberOfCompletedTours: 0,
		RequestSent:            false,
		XP:                     0,
		IsFirstPurchased:       false,
	}
	newProfile, err := service.ProfileRepository.Save(profile)
	if err != nil {
		return nil, err
	}

	return service.TokenGenerator.GenerateAccessToken(newUser, newProfile.Id)
}

func (service *AuthenticationService) Login(credentials *dto.CredentialsDto) (*dto.AuthenticationTokensDto, error) {
	user, err := service.UserRepository.GetActiveByName(credentials.Username)
	if err != nil {
		return nil, err
	}
	if user == nil || credentials.Password != user.Password {
		return nil, errors.New("not found")
	}

	personID, err := service.UserRepository.GetPersonId(user.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			personID = uuid.UUID{} // Ovdje stvarate prazan UUID
		} else {
			return nil, err
		}
	}

	return service.TokenGenerator.GenerateAccessToken(user, personID)
}
