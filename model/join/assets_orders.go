package join

type AssetsOrders struct {
	Id            int64  `gorm:"primaryKey;column:id" json:"id"`
	Uid           string `gorm:"column:uid" json:"uid"`
	TokenId       int64  `gorm:"column:token_id" json:"token_id"`
	Category      int64  `gorm:"column:category" json:"category"`
	Type          int64  `gorm:"column:type" json:"type"`
	Rarity        int64  `gorm:"column:rarity" json:"rarity"`
	Image         string `gorm:"column:image" json:"image"`
	Name          string `gorm:"column:name" json:"name"`
	Description   string `gorm:"column:description" json:"description"`
	Uri           string `gorm:"column:uri" json:"uri"`
	UriContent    string `gorm:"column:uri_content" json:"uri_content"`
	OriginChain   string `gorm:"column:origin_chain" json:"origin_chain"`
	BlockNumber   int64  `gorm:"column:block_number" json:"block_number"`
	MintSignature string `gorm:"column:mint_signature" json:"mint_signature"`
	IsNft         uint8  `gorm:"column:is_nft" json:"is_nft"` // 1: is nft    2:not nft

	Seller    string `gorm:"column:seller" json:"seller"`
	Buyer     string `gorm:"column:buyer" json:"buyer"`
	Signature string `gorm:"column:signature" json:"signature"`
	Status    uint8  `gorm:"column:status" json:"status"` // 1:active 2:expire 3:canceled 4:finished

	OrderID uint64 `gorm:"column:order_id" json:"order_id"` // orders id
	NftID   string `gorm:"column:nft_id" json:"nft_id"`
	Price   uint   `gorm:"column:price" json:"price"`
}

// TableName
func (a *AssetsOrders) TableName() string {
	return "assets"
}
