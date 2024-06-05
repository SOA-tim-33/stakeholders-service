package repository

import (
	"fmt"
	"os"
	"stakeholders-service/dto"
	"stakeholders-service/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JwtGenerator struct {
	key      []byte
	issuer   string
	audience string
}

func NewJwtGenerator() *JwtGenerator {
	key := []byte(os.Getenv("JWT_KEY"))
	if len(key) == 0 {
		key = []byte("explorer_secret_key")
	}
	issuer := os.Getenv("JWT_ISSUER")
	if issuer == "" {
		issuer = "explorer"
	}
	audience := os.Getenv("JWT_AUDIENCE")
	if audience == "" {
		audience = "explorer-front.com"
	}

	return &JwtGenerator{
		key:      key,
		issuer:   issuer,
		audience: audience,
	}
}

func (jwtGen *JwtGenerator) GenerateAccessToken(user *model.User, personID uuid.UUID) (*dto.AuthenticationTokensDto, error) {
	authenticationResponse := &dto.AuthenticationTokensDto{}

	claims := jwt.MapClaims{
		"jti":      uuid.New().String(),
		"id":       user.Id,
		"username": user.Username,
		"personId": personID,
		"role":     user.GetPrimaryRoleName(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtGen.key)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	authenticationResponse.Id = user.Id
	authenticationResponse.AccessToken = signedToken

	return authenticationResponse, nil
}
