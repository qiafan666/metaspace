package inner

type NftResp struct {
	Data []AssetDetail `json:"data"`
}

type AssetDetail struct {
	Nft NftDetail `json:"nft"`
}
type NftDetail struct {
	Id              string      `json:"id"`
	TokenId         string      `json:"tokenId"`
	ContractAddress string      `json:"contractAddress"`
	ActiveOrderId   interface{} `json:"activeOrderId"`
	Owner           string      `json:"owner"`
	Name            string      `json:"name"`
	Image           string      `json:"image"`
	Thumbnail       string      `json:"thumbnail"`
	Url             string      `json:"url"`
	Description     string      `json:"description"`
	Data            Data        `json:"data"`
	IssuedId        interface{} `json:"issuedId"`
	ItemId          interface{} `json:"itemId"`
	Category        string      `json:"category"`
	Subcategory     int64       `json:"subcategory,string"`
	Network         string      `json:"network"`
	ChainId         int         `json:"chainId"`
	CreatedAt       int64       `json:"createdAt"`
	UpdatedAt       int64       `json:"updatedAt"`
}

type Data struct {
	Tower Tower `json:"tower"`
}

type Tower struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Rarity      int64  `json:"rarity,string"`
}
