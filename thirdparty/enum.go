package thirdparty

import (
	"encoding/json"
)

type Uri string

const (
	UriEventNotify Uri = "/metaspace/event/notify"
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
	Symbol string `json:"symbol"`
}

type GameCurrencyResponse struct {
	Amount float64 `json:"amount"`
}
