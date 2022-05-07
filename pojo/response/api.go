package response

type CreateAuthCode struct {
	AuthCode               string `json:"auth_code"`
	ThirdPartyLoginAddress string `json:"third_party_login_address"`
	ThirdPartyLoginUrl     string `json:"third_party_login_url"`
}
