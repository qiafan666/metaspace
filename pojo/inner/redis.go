package inner

type Nonce struct {
	Address string `json:"address"`
	Nonce   string `json:"nonce"`
}

type PublicKey struct {
	Id                  uint64 `json:"id"`
	ApiKey              string `json:"api_key"`
	ThirdPartyPublicKey string `json:"third_party_public_key"`
}

type Rand struct {
	ApiKey string `json:"api_key"`
	Rand   string `json:"rand"`
}

type AuthCode struct {
	ThirdPartyPublicId string `json:"third_party_public_id"`
	AuthCode           string `json:"auth_code"`
	CallbackUrl        string `json:"callback_url"`
}

type ThirdPartyToken struct {
	ThirdPartyPublicId string `json:"third_party_public_id"`
	Token              string `json:"token"`
}

type TokenUser struct {
	ThirdPartyPublicId string `json:"third_party_public_id"`
	Token              string `json:"token"`
	UserId             uint64 `json:"user_id"`
	Email              string `json:"email"`
	Uuid               string `json:"uuid"`
}
