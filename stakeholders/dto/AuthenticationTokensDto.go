package dto

import "github.com/google/uuid"

type AuthenticationTokensDto struct {
	Id          uuid.UUID `json:"id"`
	AccessToken string    `json:"accessToken"`
}
