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

type AddAssets struct {
	BaseRequest
	BaseApiRequest
	AssetsList []Assets `json:"list"`
}

type Assets struct {
	UUID          string `json:"uuid"`
	WalletAddress string `json:"wallet_address"`
	Category      int64  `json:"category"`
	Type          int64  `json:"type"`
	Rarity        int64  `json:"rarity"`
	Image         string `json:"image"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Uri           string `json:"uri"`
	UriContent    string `json:"uri_content"`
}
