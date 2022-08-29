package join

import "time"

//assets,mint_history table
type MintHistoryAssets struct {
	ID            int64     `gorm:"primaryKey;column:id" json:"-"` // id
	WalletAddress string    `gorm:"column:wallet_address" json:"wallet_address"`
	TokenID       int64     `gorm:"column:token_id" json:"token_id"`
	OriginChain   uint8     `gorm:"column:origin_chain" json:"origin_chain"`
	Status        uint8     `gorm:"column:status" json:"status"`
	UpdatedTime   time.Time `gorm:"column:updated_time" json:"updated_time"` // update timestamp
	CreatedTime   time.Time `gorm:"column:created_time" json:"created_time"` // create timestamp

	NickName string `gorm:"column:nick_name" json:"nick_name"` // nick name
	Name     string `gorm:"column:name" json:"name"`           // name
}

// TableName
func (m *MintHistoryAssets) TableName() string {
	return "mint_history"
}
