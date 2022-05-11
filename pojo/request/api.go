package request

type BaseApiRequest struct {
	BaseThirdPartyId uint64 `json:"base_third_party_id"`
	BaseUserID       uint64 `json:"base_user_id"`
	BaseUUID         string `json:"base_uuid"`
	BaseEmail        string `json:"base_email"`
}

type CreateAuthCode struct {
	BaseRequest
	BaseApiRequest
}
