package join

import "time"

//orders orders_detail avatar table
type OrdersAvatar struct {
	Id         int64     `gorm:"column:id" json:"id"`
	Seller     string    `gorm:"column:seller" json:"seller"`
	Buyer      string    `gorm:"column:buyer" json:"buyer"`
	Signature  string    `gorm:"column:signature" json:"signature"`
	SaltNonce  string    `gorm:"column:salt_nonce" json:"salt_nonce"`
	Status     uint8     `gorm:"column:status" json:"status"` // 1:active 2:expire 3:canceled 4:finished
	TotalPrice string    `gorm:"column:total_price" json:"total_price"`
	StartTime  time.Time `gorm:"column:start_time" json:"start_time"`
	ExpireTime time.Time `gorm:"column:expire_time" json:"expire_time"`
	Price      string    `gorm:"column:price" json:"price"`
	NftID      int64     `gorm:"column:nft_id" json:"nft_id"`
	MarketType uint8     `gorm:"column:market_type" json:"market_type"` // 1:assets 2:avatar

	AssetId  int64  `gorm:"column:asset_id" json:"asset_id"`
	Owner    string `gorm:"column:owner" json:"owner"`
	AvatarID int64  `gorm:"column:avatar_id" json:"avatar_id"`
	Content  []byte `gorm:"column:content" json:"content"`
	IsMint   uint8  `gorm:"column:is_mint" json:"is_mint"`   // 1 minted 2 unminted
	IsShelf  uint8  `gorm:"column:is_shelf" json:"is_shelf"` // 1:shelf  2:not shelf
}

// TableName
func (m *OrdersAvatar) TableName() string {
	return "orders"
}
