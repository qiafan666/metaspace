package request

import (
	"context"
	"github.com/blockfishio/metaspace-backend/common"
)

type BaseRequest struct {
	Ctx     context.Context
	UUID    string `json:"uuid"`
	Account string `json:"account"`
}

type UserLogin struct {
	BaseRequest
	Account  string           `json:"account" validate:"required,max=192,min=6"`
	Password string           `json:"password" validate:"required,max=192,min=8"`
	Type     common.LoginType `json:"type" validate:"required,max=1"`
}

type RegisterUser struct {
	BaseRequest
	Email          string `json:"email" validate:"required,email,max=192,min=6"`
	Password       string `json:"password" validate:"required,max=192,min=8"`
	RepeatPassword string `json:"repeat_password" validate:"required,max=192,eqfield=Password"`
}

type PasswordUpdate struct {
	BaseRequest
	OldPassword       string `json:"old_password" validate:"required,max=192,min=8"`
	NewPassword       string `json:"new_password" validate:"required,max=192,min=8"`
	RepeatNewPassword string `json:"repeat_new_password" validate:"required,max=192,eqfield=NewPassword"`
}

type GetNonce struct {
	BaseRequest
	Address string `json:"address" validate:"required,eth_addr"`
}
