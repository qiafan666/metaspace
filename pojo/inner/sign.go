package inner

type SignRequest struct {
	ApiKey    string `json:"api_key"`
	Timestamp int64  `json:"timestamp"`
	Rand      string `json:"rand"`
	Uri       string `json:"uri"`
	Parameter string `json:"parameter"`
}

type SignResponse struct {
	Sign string `json:"sign"`
}

type VerifySignRequest struct {
	Sign      [32]byte `json:"sign"`
	ApiKey    string   `json:"api_key"`
	Timestamp int64    `json:"timestamp"`
	Rand      string   `json:"rand"`
	Uri       string   `json:"uri"`
	Parameter string   `json:"parameter"`
}
type VerifySignResponse struct {
	ThirdPartyId uint64 `json:"third_party_id"`
	Flag         bool   `json:"flag"`
}
