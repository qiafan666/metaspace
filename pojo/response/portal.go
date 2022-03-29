package response

type UserLogin struct {
	JwtToken string `json:"jwt_token"`
}

type RegisterUser struct {
}

type PasswordUpdate struct {
}

type GetNonce struct {
	Nonce string `json:"nonce"`
}

type GetGameAssets struct {
	AssetNum int         `json:"asset_number"`
	Assets   []AssetBody `json:"assets"`
}

type AssetBody struct {
	IsNft           bool   `json:"is_nft"`
	TokenId         string `json:"token_id"`
	ContractAddress string `json:"contract_address"`
	Name            string `json:"name"`
	Image           string `json:"image"`
	Description     string `json:"description"`
}

type SubscribeNewsletterEmail struct {
}
