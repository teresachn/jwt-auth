package payload

import (
	enums "tchen/jwt-auth/common/enums"
)

type JwtResponse struct {
	enums.ResponseBase
	AccessToken    string      `json:"accessToken"`
	TokenType      string      `json:"tokenType"`
	ExpiresIn      string      `json:"expiresIn"`
	AdditionalInfo interface{} `json:"additionalInfo"`
}
