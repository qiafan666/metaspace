package inner

type LogWriteRequest struct {
	Uri          string `json:"uri"`
	UserID       uint64 `json:"user_id"`
	ThirdPartyID uint64 `json:"third_party_id"`
	Parameter    string `json:"parameter"`
}

type LogWriteResponse struct {
}
