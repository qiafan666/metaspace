package inner

import "time"

type Nonce struct {
	Address string `json:"address"`
	Nonce   string `json:"nonce"`
}

type PublicKey struct {
	Id                  uint64 `json:"id"`
	ApiKey              string `json:"api_key"`
	ThirdPartyPublicKey string `json:"third_party_public_key"`
}

type Rand struct {
	Rand string `json:"rand"`
}

type AuthCode struct {
	ThirdPartyPublicId string `json:"third_party_public_id"`
	AuthCode           string `json:"auth_code"`
	CallbackUrl        string `json:"callback_url"`
}

type UserToken struct {
	Token  string `json:"token"`
	UserId string `json:"user_id"`
}

type TokenUser struct {
	ThirdPartyPublicId string `json:"third_party_public_id"`
	Token              string `json:"token"`
	UserId             uint64 `json:"user_id"`
	Email              string `json:"email"`
	Uuid               string `json:"uuid"`
}

type User struct {
	UserId        uint64 `json:"user_id"`
	WalletAddress string `json:"wallet_address"`
}

type RawMessage struct {
	RawMessage string    `json:"raw_message"`
	StartTime  time.Time `json:"start_time"`
	ExpireTime time.Time `json:"expire_time"`
}

type ExchangePrice struct {
	Quote      string    `json:"quote"`
	Base       string    `json:"base"`
	Price      float64   `json:"price"`
	ExpireTime time.Time `json:"expire_time"`
}
