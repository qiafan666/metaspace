package inner

type Nonce struct {
	Address string `json:"address"`
	Nonce   string `json:"nonce"`
}

type PublicKey struct {
	ApiKey              string `json:"api_key"`
	ThirdPartyPublicKey string `json:"third_party_public_key"`
}

type Rand struct {
	ApiKey string `json:"api_key"`
	Rand   string `json:"rand"`
}

type AuthCode struct {
	ApiKey string `json:"api_key"`
	Uuid   string `json:"uuid"`
}
