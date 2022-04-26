package join

//orders orders_detail assets table
type OrdersDetail struct {
	Id          int64  `gorm:"column:id" json:"id"`
	Seller      string `gorm:"column:seller" json:"seller"`
	Buyer       string `gorm:"column:buyer" json:"buyer"`
	Signature   string `gorm:"column:signature" json:"signature"`
	Status      uint8  `gorm:"column:status" json:"status"` // 1:active 2:expire 3:canceled 4:finished
	NftID       string `gorm:"column:nft_id" json:"nft_id"`
	Category    int64  `gorm:"column:category" json:"category"`
	Type        int64  `gorm:"column:type" json:"type"`
	Rarity      int64  `gorm:"column:rarity" json:"rarity"`
	Image       string `gorm:"column:image" json:"image"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}

// TableName
func (m *OrdersDetail) TableName() string {
	return "orders"
}
