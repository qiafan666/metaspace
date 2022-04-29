package request

import (
	"github.com/blockfishio/metaspace-backend/common"
)

type BasePortalRequest struct {
	BaseUserID uint64 `json:"base_user_id"`
	BaseUUID   string `json:"base_uuid"`
	BaseEmail  string `json:"base_email"`
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
}

type SubscribeNewsletterEmail struct {
	BaseRequest
	BasePortalRequest
	Email string `json:"email" validate:"required,max=192,email"`
}
type TowerStats struct {
	BaseRequest
	BasePortalRequest
	Id string `json:"id" validate:"required,max=8"`
}

type Sign struct {
	BaseRequest
	BasePortalRequest
	TokenId string `json:"token_id"  validate:"required,max=8"`
}

type ShelfSign struct {
	BaseRequest
	BasePortalRequest
	AssetId      string `json:"asset_id" validate:"required,max=192"`
	PaymentErc20 string `json:"payment_erc20" validate:"required,max=192"`
	Price        string `json:"price" validate:"required,max=192"`
}

type SellShelf struct {
	BaseRequest
	BasePortalRequest
	SignedMessage string `json:"signed_message" validate:"required,max=192"`
	RawMessage    string `json:"raw_message" validate:"required,max=192"`
	ItemId        string `json:"item_id" validate:"required,max=192"`
}

type Orders struct {
	BaseRequest
	BasePortalRequest
	Status uint8 `json:"status,string" validate:"required,max=9"`
}

type OrderCancel struct {
	BaseRequest
	BasePortalRequest
	OrderId uint64 `json:"order_id,string" validate:"required,max=192"`
}
