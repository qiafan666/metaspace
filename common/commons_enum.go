package common

import (
	"github.com/jau1jz/cornus/commons"
	"github.com/jau1jz/cornus/config"
	"time"
)

var DebugFlag bool

func init() {
	if config.SC.SConfigure.Profile == "dev" {
		DebugFlag = true
	}
}

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
	EmailAlreadyExists           = 100009
	AssetsNotExist               = 100010
	OrderAlreadyCancel           = 100011
	OrdersNotExist               = 100012
	IdentityError                = 100013
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
	EmailAlreadyExists:           "subscription email already exists",
	AssetsNotExist:               "assets is not exists",
	OrderAlreadyCancel:           "order already cancel",
	OrdersNotExist:               "orders is not exists",
	IdentityError:                "Identity does not match current login",
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

const (
	SignGrpc_CONNECT_BEFORE = 1
	SignGrpc_CONNECTING     = 2
	SignGrpc_CONNECTED      = 3
)

const GrpcTimeoutInSec = 5 * time.Second

const (
	OrderStatusActive = 1
	OrderStatusExpire = 2
	OrderStatusCancel = 3
	OrderStatusFinish = 4
)
