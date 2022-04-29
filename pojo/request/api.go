package request

import (
	"github.com/blockfishio/metaspace-backend/common"
)

type BaseApiRequest struct {
	BaseThirdPartyId uint64 `json:"base_third_party_id"`
	BaseUserID       uint64 `json:"base_user_id"`
	BaseUUID         string `json:"base_uuid"`
	BaseEmail        string `json:"base_email"`
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
