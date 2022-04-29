package request

import (
	"context"
	"github.com/blockfishio/metaspace-backend/common"
)

type BaseApiRequest struct {
	Ctx       context.Context `json:"ctx"`
	BaseUUID  string          `json:"base_uuid"`
	BaseEmail string          `json:"base_email"`
	Language  string          `json:"language"`
	ApiKey    string          `json:"api_key"`
}

type CreateAuthCode struct {
	BaseApiRequest
}

type ThirdPartyLogin struct {
	BaseApiRequest
	AuthCode string           `json:"auth_code" validate:"required,max=192"`
	Account  string           `json:"account" validate:"required,max=192,min=6"`
	Password string           `json:"password" validate:"required,max=192,min=8"`
	Type     common.LoginType `json:"type" validate:"required,max=2"`
}
