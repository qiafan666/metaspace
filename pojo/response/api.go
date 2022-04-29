package response

type CreateAuthCode struct {
	AuthCode string `json:"auth_code"`
}

type ThirdPartyLogin struct {
	JwtToken string `json:"jwt_token"`
}
