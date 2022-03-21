package inner

type NftResp struct {
	Data []AssetDetail `json:"data"`
}

type AssetDetail struct {
	Nft NftDetail `json:"nft"`
}
type NftDetail struct {
	TokenId         string `json:"tokenId"`
	ContractAddress string `json:"contractAddress"`
	Owner           string `json:"owner"`
	Name            string `json:"name"`
	Image           string `json:"image"`
	Description     string `json:"description"`
	UpdatedAt       int64  `json:"updatedAt"`
	CreatedAt       int64  `json:"createdAt"`
}
