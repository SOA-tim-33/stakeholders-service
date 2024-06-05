package handler

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"stakeholders-service/mapper"
	"stakeholders-service/service"

	"stakeholders-service/proto/stakeholder"

	"github.com/google/uuid"
)

type AuthenticationHandler struct {
	stakeholder.UnimplementedStakeholderServiceServer
	AuthenticationService *service.AuthenticationService
}

func (handler *AuthenticationHandler) RegisterTourist(ctx context.Context, request *stakeholder.RegisterTouristRequest) (*stakeholder.RegisterTouristResponse, error) {

	account := mapper.MapToModel(request)

	account.Pasysword = ToSHA256(account.Pasysword)
	token := uuid.New().String()

	result, err := handler.AuthenticationService.RegisterTourist(account, token)
	if err != nil {
		return nil, err
	}

	response := mapper.MapToProtoAuthenticationTokensDto(result)
	return response, nil

}

func (handler *AuthenticationHandler) Login(ctx context.Context, request *stakeholder.LoginRequest) (*stakeholder.RegisterTouristResponse, error) {

	credentials := mapper.MapToModelLogin(request)

	credentials.Password = ToSHA256(credentials.Password)

	result, err := handler.AuthenticationService.Login(credentials)
	if err != nil {
		return nil, nil
	}

	response := mapper.MapToProtoAuthenticationTokensDto(result)
	return response, nil

}

func ToSHA256(s string) string {
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(s))
	hashedBytes := sha256Hash.Sum(nil)

	hashedString := hex.EncodeToString(hashedBytes)
	return hashedString
}
