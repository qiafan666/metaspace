package common

import (
	"github.com/qiafan666/quickweb/commons"
	"github.com/qiafan666/quickweb/config"
)

var DebugFlag bool

func init() {
	if config.SC.SConfigure.Profile == "dev" {
		DebugFlag = true
	}
}

// define the error code
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
	ThirdPartySignError          = 100014
	VerifyThirdPartySignError    = 100015
	VerifyThirdPartySignTimeOut  = 100016
	FrequentVerifyThirdPartySign = 100017
	AuthCodeAlreadyExpired       = 100018
	OrdersIsShelf                = 100019
	UsedSignature                = 100020
	WalletError                  = 100021
	HistoryError                 = 100022
	GameCurrencyError            = 100023
	ChainNetError                = 100024
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
	AssetsNotExist:               "Asset doesn't exist",
	OrderAlreadyCancel:           "Order is already cancelled.",
	OrdersNotExist:               "Order doesn't exist",
	IdentityError:                "Identity check failed",
	ThirdPartySignError:          "Third_Party Sign failed",
	VerifyThirdPartySignError:    "Verify Third_Party Sign failed",
	VerifyThirdPartySignTimeOut:  "Verify Third_Party Sign timeout",
	FrequentVerifyThirdPartySign: "Frequent Verify Third_Party Sign",
	AuthCodeAlreadyExpired:       "AuthCode is already expired",
	OrdersIsShelf:                "Orders is already shelf",
	UsedSignature:                "Signatures is already used",
	WalletError:                  "Inconsistent wallet addresses",
	HistoryError:                 "History type doesn't exist",
	GameCurrencyError:            "Failed to get game data",
	ChainNetError:                "Current network is not supported",
}

// login type
type LoginType uint8

const (
	LoginTypeEmail LoginType = iota + 1
	LoginTypeWallet
)

// redis key
const (
	RawMessage = "user/raw_message/%s"
)

// ctx value enum
const (
	BaseRequest       = "base_request"
	BasePortalRequest = "base_portal_request"
	BaseApiRequest    = "base_api_request"
)
