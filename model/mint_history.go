package model

import (
	"time"
)

/******sql******
CREATE TABLE `mint_history` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `wallet_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `token_id` bigint NOT NULL,
  `origin_chain` bigint unsigned NOT NULL,
  `status` tinyint unsigned NOT NULL COMMENT '1:mint',
  `updated_time` timestamp(3) NOT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'update timestamp',
  `created_time` timestamp(3) NOT NULL COMMENT 'create timestamp',
  PRIMARY KEY (`id`),
  KEY `token_id` (`token_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=176 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// MintHistory [...]
type MintHistory struct {
	ID            int64     `gorm:"primaryKey;column:id" json:"-"` // id
	WalletAddress string    `gorm:"column:wallet_address" json:"wallet_address"`
	TokenID       int64     `gorm:"column:token_id" json:"token_id"`
	OriginChain   uint64    `gorm:"column:origin_chain" json:"origin_chain"`
	Status        uint8     `gorm:"column:status" json:"status"`             // 1:mint
	UpdatedTime   time.Time `gorm:"column:updated_time" json:"updated_time"` // update timestamp
	CreatedTime   time.Time `gorm:"column:created_time" json:"created_time"` // create timestamp
}

// TableName get sql table name.获取数据库表名
func (m *MintHistory) TableName() string {
	return "mint_history"
}

// MintHistoryColumns get sql column name.获取数据库列名
var MintHistoryColumns = struct {
	ID            string
	WalletAddress string
	TokenID       string
	OriginChain   string
	Status        string
	UpdatedTime   string
	CreatedTime   string
}{
	ID:            "id",
	WalletAddress: "wallet_address",
	TokenID:       "token_id",
	OriginChain:   "origin_chain",
	Status:        "status",
	UpdatedTime:   "updated_time",
	CreatedTime:   "created_time",
}
