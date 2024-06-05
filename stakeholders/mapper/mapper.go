package mapper

import (
	"stakeholders-service/dto"
	"stakeholders-service/proto/stakeholder"
)

func MapToModel(touristRequest *stakeholder.RegisterTouristRequest) *dto.AccountRegistrationDto {
	proto := &dto.AccountRegistrationDto{
		Id:             touristRequest.Id,
		Name:           touristRequest.Name,
		Surname:        touristRequest.Surname,
		Email:          touristRequest.Email,
		Username:       touristRequest.Username,
		Pasysword:      touristRequest.Password,
		ProfilePicture: touristRequest.ProfilePicture,
		Biography:      touristRequest.Biography,
		Motto:          touristRequest.Motto,
	}

	return proto
}

func MapToProtoAuthenticationTokensDto(req *dto.AuthenticationTokensDto) *stakeholder.RegisterTouristResponse {
	proto := &stakeholder.RegisterTouristResponse{
		Id:          req.Id.String(),
		AccessToken: req.AccessToken,
	}

	return proto
}

func MapToModelLogin(touristRequest *stakeholder.LoginRequest) *dto.CredentialsDto {
	proto := &dto.CredentialsDto{
		Username: touristRequest.Username,
		Password: touristRequest.Password,
	}

	return proto
}
