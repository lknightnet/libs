package genToken

type GenerateTokenRequest struct {
	ClientUUID string `json:"client_uuid"`
}

type GetAccessTokenRequest struct {
	AccessToken string `json:"access_token"`
}

type GetRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}
