package response

import (
	"time"
)

type BasePagination struct {
	Total        int64 `json:"total"`
	CurrentPage  int   `json:"currentPage"`
	PrePageCount int   `json:"prePageCount"`
}

type ThirdPartyLogin struct {
	Token         string    `json:"token"`
	Email         string    `json:"email"`
	Uuid          string    `json:"uuid"`
	WalletAddress string    `json:"wallet_address"`
	Url           string    `json:"url,omitempty"`
	ExpireTime    time.Time `json:"expire_time"`
}

type UserLogin struct {
	JwtToken      string `json:"jwt_token"`
	UserName      string `json:"user_name"`
	AvatarAddress string `json:"avatar_address"`
}

type RegisterUser struct {
}

type PasswordUpdate struct {
}

type GetNonce struct {
	Nonce string `json:"nonce"`
}

type GetGameAssets struct {
	BasePagination
	Assets []AssetBody `json:"assets"`
}

type AssetBody struct {
	AssetId         int64     `json:"asset_id"`
	IsNft           uint8     `json:"is_nft"`
	TokenId         int64     `json:"token_id"`
	ContrainChain   uint64    `json:"contract_chain"`
	ContractAddress string    `json:"contract_address"`
	Name            string    `json:"name"`
	IndexID         uint64    `json:"index_id"`
	NickName        string    `json:"nick_name"`
	Image           string    `json:"image"`
	Description     string    `json:"description"`
	Category        string    `json:"category"`
	CategoryId      int64     `json:"category_id"`
	Rarity          string    `json:"rarity"`
	RarityId        int64     `json:"rarity_id"`
	MintSignature   string    `json:"mint_signature"`
	Subcategory     string    `json:"subcategory"`
	SubcategoryId   int64     `json:"subcategory_id"`
	Status          uint8     `json:"status"`
	Price           string    `json:"price"`
	OrderId         uint64    `json:"order_id"`
	ExpireTime      time.Time `json:"expire_time"`
	Signature       string    `json:"signature"`
	SaltNonce       string    `json:"salt_nonce"`
	StartTime       time.Time `json:"start_time"`
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
	SignMessage     string `json:"sign_message"`
	ContractAddress string `json:"contract_address"`
}

type ShelfSign struct {
	SignMessage string `json:"sign_message"`
	SaltNonce   string `json:"salt_nonce"`
}

type SellShelf struct {
	RawMessage  string `json:"raw_message"`
	SignMessage string `json:"sign_message"`
}

type Orders struct {
	BasePagination
	Data []OrdersDetail `json:"orders_list"`
}

type OrdersDetail struct {
	AssetId         int64     `json:"asset_id"`
	Id              int64     `json:"id"`
	Seller          string    `json:"seller"`
	Buyer           string    `json:"buyer"`
	Signature       string    `json:"signature"`
	SaltNonce       int64     `json:"salt_nonce"`
	Status          uint8     `json:"status"` // 1:active 2:expire 3:canceled 4:finished
	NftID           int64     `json:"token_id"`
	Category        int64     `json:"category_id"`
	Type            int64     `json:"type"`
	Rarity          int64     `json:"rarity_id"`
	Image           string    `json:"image"`
	Name            string    `json:"name"`
	IndexID         uint64    `json:"index_id"`
	NickName        string    `json:"nick_name"`
	Description     string    `json:"description"`
	TotalPrice      string    `json:"total_price"`
	Price           string    `json:"price"`
	ContractChain   uint64    `json:"contract_chain"`
	StartTime       time.Time `json:"start_time"`
	ExpireTime      time.Time `json:"expire_time"`
	ContractAddress string    `json:"contract_address"`
	GroupName       string    `json:"group_name"`
}

type OrderCancel struct {
	OrderId uint64 `json:"order_id"`
}

type UserUpdate struct {
}

type UserHistory struct {
	BasePagination
	Data []HistoryList `json:"history_list"`
}

type HistoryList struct {
	WalletAddress string    `json:"wallet_address"`
	TokenID       int64     `json:"token_id"`
	Price         string    `json:"price"`
	Unit          string    `json:"unit"`
	Status        uint8     `json:"status"`                            // 1:上架 2:下架 3:买 4:卖]
	CreatedTime   time.Time `json:"created_time"`                      // create timestamp
	Name          string    `gorm:"column:name" json:"name"`           // name
	NickName      string    `gorm:"column:nick_name" json:"nick_name"` // nick name
	IndexID       uint64    `gorm:"column:index_id" json:"index_id"`
}

type ExchangePrice struct {
	Price float64 `json:"price"`
}

type AssetDetail struct {
	AssetId         int64     `json:"asset_id"`
	WalletAddress   string    `json:"seller"`
	IsNft           uint8     `json:"is_nft"`
	TokenId         int64     `json:"token_id"`
	ContrainChain   uint64    `json:"contract_chain"`
	ContractAddress string    `json:"contract_address"`
	Name            string    `json:"name"`
	IndexID         uint64    `json:"index_id"`
	NickName        string    `json:"nick_name"`
	Image           string    `json:"image"`
	Description     string    `json:"description"`
	Category        string    `json:"category"`
	CategoryId      int64     `json:"category_id"`
	Rarity          string    `json:"rarity"`
	RarityId        int64     `json:"rarity_id"`
	MintSignature   string    `json:"mint_signature"`
	Subcategory     string    `json:"subcategory"`
	SubcategoryId   int64     `json:"subcategory_id"`
	Status          uint8     `json:"status"`
	Price           string    `json:"price"`
	OrderId         uint64    `json:"order_id"`
	ExpireTime      time.Time `json:"expire_time"`
	Signature       string    `json:"signature"`
	SaltNonce       string    `json:"salt_nonce"`
	StartTime       time.Time `json:"start_time"`
	GroupName       string    `json:"group_name"`
}

type GameCurrency struct {
	Amount int64 `json:"amount"`
}

type OrdersGroup struct {
	BasePagination
	Data []OrdersDetail `json:"orders_list"`
}

type OrdersGroupDetail struct {
	AssetId         int64     `json:"asset_id"`
	Id              int64     `json:"id"`
	Seller          string    `json:"seller"`
	Buyer           string    `json:"buyer"`
	Signature       string    `json:"signature"`
	SaltNonce       int64     `json:"salt_nonce"`
	Status          uint8     `json:"status"` // 1:active 2:expire 3:canceled 4:finished
	NftID           int64     `json:"token_id"`
	Category        int64     `json:"category_id"`
	Type            int64     `json:"type"`
	Rarity          int64     `json:"rarity_id"`
	Image           string    `json:"image"`
	Name            string    `json:"name"`
	IndexID         uint64    `json:"index_id"`
	NickName        string    `json:"nick_name"`
	Description     string    `json:"description"`
	TotalPrice      string    `json:"total_price"`
	Price           string    `json:"price"`
	ContractChain   uint64    `json:"contract_chain"`
	StartTime       time.Time `json:"start_time"`
	ExpireTime      time.Time `json:"expire_time"`
	ContractAddress string    `json:"contract_address"`
}

type SendCode struct {
}

type PaperMint struct {
	SdkClientSecret string `json:"sdkClientSecret"`
}

type OrdersOfficial struct {
	BasePagination
	Data []OrdersDetail `json:"orders_list"`
}

type PaperTransaction struct {
	SdkClientSecret string `json:"sdkClientSecret"`
}

type Avatar struct {
	BasePagination
	Data []AvatarBody `json:"avatar_list"`
}

type AvatarBody struct {
	Id            int64     `json:"asset_id"`
	Owner         string    `json:"seller"`
	AvatarID      int64     `json:"token_id"`
	Content       string    `json:"content"`
	IsShelf       uint8     `json:"is_shelf"` // 1:shelf  2:not shelf
	Status        uint8     `json:"status"`
	Price         string    `json:"price"`
	OrderId       uint64    `json:"order_id"`
	ExpireTime    time.Time `json:"expire_time"`
	Signature     string    `json:"signature"`
	SaltNonce     string    `json:"salt_nonce"`
	StartTime     time.Time `json:"start_time"`
	ContractChain uint64    `json:"contract_chain"`
}

type OrderAvatar struct {
	BasePagination
	Data []AvatarDetail `json:"avatar_list"`
}

type AvatarDetail struct {
	Id            int64     `json:"order_id"`
	AssetId       int64     `json:"asset_id"`
	Owner         string    `json:"seller"`
	AvatarID      int64     `json:"token_id"`
	Content       string    `json:"content"`
	Price         string    `json:"price"`
	Status        uint8     `json:"status"`
	Signature     string    `json:"signature"`
	SaltNonce     string    `json:"salt_nonce"`
	StartTime     time.Time `json:"start_time"`
	ExpireTime    time.Time `json:"expire_time"`
	ContractChain uint64    `json:"contract_chain"`
	IsNft         uint8     `json:"is_nft"`
}

type Test struct {
}
