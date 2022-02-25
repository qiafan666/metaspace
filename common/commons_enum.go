package common

import (
	"github.com/jau1jz/cornus/commons"
)

//define the error code
const (
	PasswordOrAccountError      = 100001
	AccountAlreadyExists        = 100002
	OldPasswordNotEqual         = 100003
	OldPasswordEqualNewPassword = 100004
)

// CodeMsg local code and msg

var CodeMsg = map[commons.ResponseCode]string{
	PasswordOrAccountError:      "account or password error.",
	AccountAlreadyExists:        "account already exists.",
	OldPasswordNotEqual:         "old password not equal",
	OldPasswordEqualNewPassword: "old password equal new password",
}

// login type
type LoginType uint8

const (
	LoginTypeEmail LoginType = iota + 1
)
