package payload

type JwtRequest struct {
	GrantType      string      `json:"grantType"`
	AdditionalInfo interface{} `json:"additionalInfo"`
}
