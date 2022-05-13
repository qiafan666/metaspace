package join

import "time"

//orders,orders_detail table
type OrdersStatus struct {
	Id         int64     `gorm:"column:id" json:"id"`
	Seller     string    `gorm:"column:seller" json:"seller"`
	Buyer      string    `gorm:"column:buyer" json:"buyer"`
	Signature  string    `gorm:"column:signature" json:"signature"`
	Status     uint8     `gorm:"column:status" json:"status"` // 1:active 2:expire 3:canceled 4:finished
	ExpireTime time.Time `gorm:"column:expire_time" json:"expire_time"`
	NftID      string    `gorm:"column:nft_id" json:"nft_id"`
}

// TableName
func (m *OrdersStatus) TableName() string {
	return "orders"
}
