package request

import (
	"github.com/blockfishio/metaspace-backend/common"
	"time"
)

type BasePortalRequest struct {
	BaseUserID uint64 `json:"base_user_id"`
	BaseUUID   string `json:"base_uuid"`
	BaseEmail  string `json:"base_email"`
	BaseWallet string `json:"base_wallet"`
}

type ThirdPartyLogin struct {
	BaseRequest
	AuthCode string           `json:"auth_code" validate:"required,max=192"`
	Account  string           `json:"account" validate:"required,max=192,min=6"`
	Password string           `json:"password" validate:"required,max=192,min=8"`
	Type     common.LoginType `json:"type" validate:"required,max=2"`
}

type UserLogin struct {
	BaseRequest
	BasePortalRequest
	Account  string           `json:"account" validate:"required,max=192,min=6"`
	Password string           `json:"password" validate:"required,max=192,min=8"`
	Type     common.LoginType `json:"type" validate:"required,max=2"`
}

type RegisterUser struct {
	BaseRequest
	BasePortalRequest
	Email          string `json:"email" validate:"required,email,max=192,min=6"`
	Password       string `json:"password" validate:"required,max=192,min=8"`
	RepeatPassword string `json:"repeat_password" validate:"required,max=192,eqfield=Password"`
}

type PasswordUpdate struct {
	BaseRequest
	BasePortalRequest
	OldPassword       string `json:"old_password" validate:"required,max=192,min=8"`
	NewPassword       string `json:"new_password" validate:"required,max=192,min=8"`
	RepeatNewPassword string `json:"repeat_new_password" validate:"required,max=192,eqfield=NewPassword"`
}

type GetNonce struct {
	BaseRequest
	BasePortalRequest
	Address string `json:"address" validate:"required,eth_addr"`
}

type GetGameAssets struct {
	BaseRequest
	BasePortalRequest
	BasePagination
	Category *int  `json:"category" validate:""`
	Rarity   *int  `json:"rarity" validate:""`
	IsNft    *int  `json:"is_nft" validate:""`
	IsShelf  int   `json:"is_sale" validate:"max=2"`
	ChainId  uint8 `json:"chain_id"`
}

type SubscribeNewsletterEmail struct {
	BaseRequest
	BasePortalRequest
	Email string `json:"email" validate:"required,max=192,email"`
}
type TowerStats struct {
	BaseRequest
	BasePortalRequest
	Id string `json:"id" validate:"required,max=192"`
}

type Sign struct {
	BaseRequest
	BasePortalRequest
	Chain   uint8 `json:"chain" validate:"required"`
	TokenId int64 `json:"token_id" validate:"required"`
}

type ShelfSign struct {
	BaseRequest
	BasePortalRequest
	AssetId      string    `json:"asset_id" validate:"required,max=192"`
	PaymentErc20 string    `json:"payment_erc20" validate:"required,max=192"`
	Price        string    `json:"price" validate:"required,max=192"`
	ExpireTime   time.Time `json:"expire_time" validate:"required"`
}

type SellShelf struct {
	BaseRequest
	BasePortalRequest
	SignedMessage string `json:"signed_message" validate:"required,max=192"`
	RawMessage    string `json:"raw_message" validate:"required,max=192"`
	ItemId        string `json:"item_id" validate:"required,max=192"`
	Price         string `json:"price" validate:"required"`
	SaltNonce     string `json:"salt_nonce" validate:"required"`
}

type Orders struct {
	BaseRequest
	BasePagination
	BasePortalRequest
	Status   uint8 `json:"status,string" validate:"max=5"`
	Category *int  `json:"category" validate:""`
	Rarity   *int  `json:"rarity" validate:""`
	//sort
	SortPrice uint   `json:"sort_price" validate:"max=2"`
	SortTime  uint   `json:"sort_time" validate:"max=2"`
	Search    string `json:"search" validate:"max=192"`

	ChainId uint8 `json:"chain_id"`
}

type OrderCancel struct {
	BaseRequest
	BasePortalRequest
	OrderId uint64 `json:"order_id,string" validate:"required"`
}

type UserUpdate struct {
	BaseRequest
	BasePortalRequest
	UserName      string `json:"user_name"`
	AvatarAddress string `json:"avatar_address"`
}

type UserHistory struct {
	BaseRequest
	BasePortalRequest
	BasePagination
	Type              uint8     `json:"type" validate:"required"` //1:transaction history  //2:mint history //3:Listing history
	FilterTransaction uint8     `json:"filter_transaction"`
	FilterTime        time.Time `json:"filter_time"`
	ChainId           uint8     `json:"chain_id"`
}

type ExchangePrice struct {
	BaseRequest
	BasePortalRequest
	Quote string `json:"quote" validate:"required"`
	Base  string `json:"base"  validate:"required"`
}

type AssetDetail struct {
	BaseRequest
	BasePortalRequest
	AssetId int64  `json:"asset_id"`
	ChainId string `json:"chain_id"`
	TokenId string `json:"token_id"`
}

type GameCurrency struct {
	BaseRequest
	BasePortalRequest
	Symbol string `json:"symbol" validate:"required"`
}

type OrdersGroup struct {
	BaseRequest
	BasePagination
	BasePortalRequest
	Status   uint8 `json:"status,string" validate:"max=5"`
	Category *int  `json:"category" validate:""`
	Rarity   *int  `json:"rarity" validate:""`
	//sort
	SortPrice uint   `json:"sort_price" validate:"max=2"`
	SortTime  uint   `json:"sort_time" validate:"max=2"`
	Search    string `json:"search" validate:"max=192"`

	ChainId uint8 `json:"chain_id"`
}

type OrdersGroupDetail struct {
	BaseRequest
	GroupName string `json:"group_name" validate:"required"`
	SortPrice uint   `json:"sort_price" validate:"max=2"`
	ChainId   uint8  `json:"chain_id"`
}

type SendCode struct {
	BaseRequest
	Email string `json:"email" validate:"required,email"`
}

type PaperMint struct {
	BaseRequest
	TokenId       int64  `json:"token_id"`
	ChainId       uint8  `json:"chain_id"`
	Email         string `json:"email"`
	WalletAddress string `json:"wallet_address"`
}

type PaperMintRequest struct {
	Quantity int `json:"quantity"`
	Metadata struct {
	} `json:"metadata"`
	ExpiresInMinutes      int    `json:"expiresInMinutes"`
	UsePaperKey           bool   `json:"usePaperKey"`
	HideApplePayGooglePay bool   `json:"hideApplePayGooglePay"`
	ContractId            string `json:"contractId"`
	WalletAddress         string `json:"walletAddress"`
	Email                 string `json:"email"`
	MintMethod            struct {
		Name string `json:"name"`
		Args struct {
			NftAddress  string `json:"nftAddress"`
			UserAddress string `json:"userAddress"`
			TokenId     int64  `json:"tokenId"`
			Category    int64  `json:"category"`
			Subcategory int64  `json:"subcategory"`
			Rarity      int64  `json:"rarity"`
			Signature   string `json:"signature"`
		} `json:"args"`
		Payment struct {
			Value    string `json:"value"`
			Currency string `json:"currency"`
		} `json:"payment"`
	} `json:"mintMethod"`
}
