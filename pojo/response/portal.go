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
	ContrainChain   string `json:"contract_chain"`
	ContractAddress string `json:"contract_address"`
	Name            string `json:"name"`
	Image           string `json:"image"`
	Description     string `json:"description"`
	Category        string `json:"category"`
	CategoryId      int    `json:"category_id"`
	Rarity          string `json:"rarity"`
	RarityId        int    `json:"rarity_id"`
	MintSignature   string `json:"mint_signature"`
}
