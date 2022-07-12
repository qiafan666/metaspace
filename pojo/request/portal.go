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
	Category *int `json:"category"`
	Rarity   *int `json:"rarity"`
	IsNft    *int `json:"is_nft"`
	IsSale   int  `json:"is_sale"`
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
	TokenId int64 `json:"token_id"  validate:"required"`
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
	Status   uint8 `json:"status,string"`
	Category *int  `json:"category"`
	Rarity   *int  `json:"rarity"`
	//sort
	SortPrice uint `json:"sort_price"`
	SortTime  uint `json:"sort_time"`
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
	AssetId int64 `json:"asset_id" validate:"required"`
}
