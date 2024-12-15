package genToken

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type AccessTokenResponse struct {
	Token      string `json:"token"`
	Subject    string `json:"subject"`
	Expiration int64  `json:"expiration"`
	IssuedAt   int64  `json:"issued_at"`
}

type RefreshTokenResponse struct {
	Token      string `json:"token"`
	Subject    string `json:"subject"`
	Expiration int64  `json:"expiration"`
	IssuedAt   int64  `json:"issued_at"`
}
