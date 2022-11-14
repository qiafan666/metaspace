package request

import (
	"github.com/qiafan666/fundametality/common"
	"time"
)

type BasePortalRequest struct {
	BaseUserID uint64 `json:"base_user_id"`
	BaseUUID   string `json:"base_uuid"`
	BaseEmail  string `json:"base_email"`
	BaseWallet string `json:"base_wallet"`
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
