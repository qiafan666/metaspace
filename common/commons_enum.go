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
	ThirdPartySignError          = 100014
	VerifyThirdPartySignError    = 100015
	VerifyThirdPartySignTimeOut  = 100016
	FrequentVerifyThirdPartySign = 100017
	AuthCodeAlreadyExpired       = 100018
	OrdersIsShelf                = 100019
	UsedSignature                = 100020
	WalletError                  = 100021
	HistoryError                 = 100022
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
}

// login type
type LoginType uint8

const (
	LoginTypeEmail LoginType = iota + 1
	LoginTypeWallet
)

// redis key
const (
	UserNonce           = "user/nonce/%s"
	ThirdPartyPublicKey = "third_party/publicKey/%s"
	ThirdPartyRand      = "third_party/rand/%s"
	ThirdPartyAuthCode  = "third_party/auth_code/%s"
	ThirdPartyUserToken = "third_party/user_token/%s/%s"
	ThirdPartyTokenUser = "third_party/token_user/%s"
	RawMessage          = "user/raw_message/%s"
)

const (
	SignGrpcConnectBefore = 1
	SignGrpcConnecting    = 2
	SignGrpcConnected     = 3
)

const GrpcTimeoutIn = 5 * time.Second

const (
	OrderStatusActive = 1
	OrderStatusExpire = 2
	OrderStatusCancel = 3
	OrderStatusFinish = 4
)

//ctx value enum
const (
	BaseRequest       = "base_request"
	BasePortalRequest = "base_portal_request"
	BaseApiRequest    = "base_api_request"
)

//url
const (
	UrlCallbackLogin = "/metaspace/callback/login"
)

//third login
const (
	BaseRequestSign          = "Sign"
	BaseRequestApiKey        = "Api-key"
	BaseRequestTimestamp     = "Timestamp"
	BaseRequestRand          = "Rand"
	BaseRequestAuthorization = "Authorization"
)

const (
	IsNft  = 1
	NotNft = 2
)

const (
	TransactionHistory = 1
	MintHistory        = 2
	ListenHistory      = 3
)

//transaction_history activity type
const (
	Shelf     = 1
	Cancel    = 2
	Sold      = 3
	Purchased = 4
)
