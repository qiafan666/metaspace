package thirdparty

import (
	"encoding/json"
)

type Uri string

const (
	UriWalletBalance Uri = "/metaspace/wallet/balance"
)

type ResponseCode int
type BaseRequest struct {
	Code ResponseCode    `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
	Time int64           `json:"time"`
}
type BaseResponse struct {
	Code ResponseCode `json:"code"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
	Time int64        `json:"time"`
}

type BaseNotifyEvent struct {
	Type      uint        `json:"type"`
	EventData interface{} `json:"event_data"`
}

type BaseThirdParty struct {
	ThirdPartyID uint64 `json:"third_party_id"`
}

//third login
const (
	HeaderSign          = "Sign"
	HeaderApiKey        = "Api-key"
	HeaderTimestamp     = "Timestamp"
	HeaderRand          = "Rand"
	HeaderAuthorization = "Authorization"
)

type GameCurrencyRequest struct {
	WallAddress string `json:"wall_address"`
	Symbol      string `json:"symbol"`
}

type GameCurrencyResponse struct {
	Amount int64 `json:"amount"`
}
