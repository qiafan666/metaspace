package request

import (
	"github.com/blockfishio/metaspace-backend/common"
)

type BaseApiRequest struct {
	ApiKey string `json:"api_key"`
}

type CreateAuthCode struct {
	BaseRequest
	BaseApiRequest
}

type ThirdPartyLogin struct {
	BaseRequest
	BaseApiRequest
	AuthCode string           `json:"auth_code" validate:"required,max=192"`
	Account  string           `json:"account" validate:"required,max=192,min=6"`
	Password string           `json:"password" validate:"required,max=192,min=8"`
	Type     common.LoginType `json:"type" validate:"required,max=2"`
}
