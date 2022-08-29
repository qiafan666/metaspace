package join

import "time"

//orders orders_detail assets table
type OrdersDetail struct {
	Id          int64     `gorm:"column:id" json:"id"`
	Seller      string    `gorm:"column:seller" json:"seller"`
	Buyer       string    `gorm:"column:buyer" json:"buyer"`
	Signature   string    `gorm:"column:signature" json:"signature"`
	SaltNonce   int64     `gorm:"column:salt_nonce" json:"salt_nonce"`
	Status      uint8     `gorm:"column:status" json:"status"` // 1:active 2:expire 3:canceled 4:finished
	TotalPrice  string    `gorm:"column:total_price" json:"total_price"`
	StartTime   time.Time `gorm:"column:start_time" json:"start_time"`
	ExpireTime  time.Time `gorm:"column:expire_time" json:"expire_time"`
	Price       string    `gorm:"column:price" json:"price"`
	NftID       int64     `gorm:"column:nft_id" json:"nft_id"`
	AssetId     int64     `gorm:"column:asset_id" json:"asset_id"`
	Category    int64     `gorm:"column:category" json:"category"`
	Type        int64     `gorm:"column:type" json:"type"`
	Rarity      int64     `gorm:"column:rarity" json:"rarity"`
	Image       string    `gorm:"column:image" json:"image"`
	Name        string    `gorm:"column:name" json:"name"`
	IndexID     uint64    `gorm:"column:index_id" json:"index_id"`
	NickName    string    `gorm:"column:nick_name" json:"nick_name"`
	Description string    `gorm:"column:description" json:"description"`
	OriginChain uint8     `gorm:"column:origin_chain" json:"origin_chain"`
	Sku         string    `gorm:"column:sku" json:"sku"`
}

// TableName
func (m *OrdersDetail) TableName() string {
	return "orders"
}
