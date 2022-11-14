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
	PasswordOrAccountError = 100001
)

// EnglishCodeMsg local code and msg

var EnglishCodeMsg = map[commons.ResponseCode]string{
	PasswordOrAccountError: "account or password error",
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
