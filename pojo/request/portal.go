package request

import (
	"github.com/qiafan666/fundametality/common"
)

type BasePortalRequest struct {
	BaseUUID  string `json:"base_uuid"`
	BaseEmail string `json:"base_email"`
}

type UserLogin struct {
	BaseRequest
	BasePortalRequest
	Account  string           `json:"account" validate:"required,max=192,min=6"`
	Password string           `json:"password" validate:"required,max=192,min=8"`
	Type     common.LoginType `json:"type" validate:"required,max=2"`
}
