package common

import (
	"github.com/jau1jz/cornus/commons"
)

//define the error code
const (
	PasswordOrAccountError       = 100001
	AccountAlreadyExists         = 100002
	OldPasswordNotEqual          = 100003
	OldPasswordEqualNewPassword  = 100004
	WalletAddressDoesNotRegister = 100005
	SignatureVerificationError   = 100006
	WalletAddressDoesNotExist    = 100007
	NonceExpireOrNull            = 100008
)

// EnglishCodeMsg local code and msg

var EnglishCodeMsg = map[commons.ResponseCode]string{
	PasswordOrAccountError:       "account or password error",
	AccountAlreadyExists:         "account already exists",
	OldPasswordNotEqual:          "old password not equal",
	OldPasswordEqualNewPassword:  "old password equal new password",
	WalletAddressDoesNotRegister: "wallet address does not register",
	SignatureVerificationError:   "signature verification error",
	WalletAddressDoesNotExist:    "wallet address does not exist",
	NonceExpireOrNull:            "nonce expire or null",
}

// login type
type LoginType uint8

const (
	LoginTypeEmail LoginType = iota + 1
	LoginTypeWallet
)

// redis key
const (
	UserNonce = "user/nonce/%s"
)
