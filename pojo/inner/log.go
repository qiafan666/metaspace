package inner

import "context"

type LogWriteRequest struct {
	Ctx          context.Context
	Uri          string `json:"uri"`
	UserID       uint64 `json:"user_id"`
	ThirdPartyID uint64 `json:"third_party_id"`
	Parameter    []byte `json:"parameter"`
}
