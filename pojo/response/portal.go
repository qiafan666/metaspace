package response

import "github.com/blockfishio/metaspace-backend/model"

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
	CategoryId      int64  `json:"category_id"`
	Rarity          string `json:"rarity"`
	RarityId        int64  `json:"rarity_id"`
	MintSignature   string `json:"mint_signature"`
	Subcategory     string `json:"subcategory"`
	SubcategoryId   int64  `json:"subcategory_id"`
}

type SubscribeNewsletterEmail struct {
}

type TowerStats struct {
	Attack      int     `json:"attack"`
	FireRate    float32 `json:"fire_rate"`
	AttackRange int     `json:"attack_range"`
	Durability  int     `json:"durability"`
}

type Sign struct {
	SignMessage string `json:"sign_message"`
}

type ShelfSign struct {
	SignMessage string `json:"sign_message"`
}

type SellShelf struct {
	Flag string `json:"flag"`
}

type Orders struct {
	Data []model.Orders `json:"data"`
}
