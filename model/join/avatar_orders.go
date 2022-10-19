package join

import "time"

//avatar,orders table
type AvatarOrders struct {
	ID       int64  `gorm:"primaryKey;column:id" json:"-"` // id
	Owner    string `gorm:"column:owner" json:"owner"`
	AvatarID int64  `gorm:"column:avatar_id" json:"avatar_id"`
	Content  []byte `gorm:"column:content" json:"content"`
	IsMint   uint8  `gorm:"column:is_mint" json:"is_mint"`   // 1 minted 2 unminted
	IsShelf  uint8  `gorm:"column:is_shelf" json:"is_shelf"` // 1:shelf  2:not shelf

	Seller     string    `gorm:"column:seller" json:"seller"`
	Buyer      string    `gorm:"column:buyer" json:"buyer"`
	Signature  string    `gorm:"column:signature" json:"signature"`
	Status     uint8     `gorm:"column:status" json:"status"` // 1:active 2:expire 3:canceled 4:finished
	SaltNonce  string    `gorm:"column:salt_nonce" json:"salt_nonce"`
	ExpireTime time.Time `gorm:"column:expire_time" json:"expire_time"`
	StartTime  time.Time `gorm:"column:start_time" json:"start_time"`

	OrderID int64  `gorm:"column:order_id" json:"order_id"` // orders id
	NftID   string `gorm:"column:nft_id" json:"nft_id"`
	Price   string `gorm:"column:price" json:"price"`
}

// TableName
func (a *AvatarOrders) TableName() string {
	return "avatar"
}
