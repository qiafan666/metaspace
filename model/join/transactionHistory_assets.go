package join

import "time"

//assets,transaction_history table
type TransactionHistoryAssets struct {
	ID            int64     `gorm:"primaryKey;column:id" json:"-"` // asset id
	WalletAddress string    `gorm:"column:wallet_address" json:"wallet_address"`
	TokenID       int64     `gorm:"column:token_id" json:"token_id"`
	Price         string    `gorm:"column:price" json:"price"`
	Unit          string    `gorm:"column:Unit" json:"unit"`
	OriginChain   uint8     `gorm:"column:origin_chain" json:"origin_chain"`
	Status        uint8     `gorm:"column:status" json:"status"`             // 1:上架 2:下架 3:买 4:卖]
	CreatedTime   time.Time `gorm:"column:created_time" json:"created_time"` // create timestamp

	NickName string `gorm:"column:nick_name" json:"nick_name"` // nick name
	Name     string `gorm:"column:name" json:"name"`           // name
	IndexID  uint64 `gorm:"column:index_id" json:"index_id"`
}

// TableName
func (t *TransactionHistoryAssets) TableName() string {
	return "transaction_history"
}
