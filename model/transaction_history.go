package model

import (
	"time"
)

/******sql******
CREATE TABLE `transaction_history` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'asset id',
  `wallet_address` varchar(255) NOT NULL,
  `token_id` bigint NOT NULL,
  `price` varchar(192) NOT NULL,
  `Unit` varchar(192) NOT NULL,
  `status` tinyint unsigned NOT NULL COMMENT '1:上架 2:下架 3:买 4:卖',
  `updated_time` datetime(3) NOT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'update timestamp',
  `created_time` datetime(3) NOT NULL COMMENT 'create timestamp',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// TransactionHistory [...]
type TransactionHistory struct {
	ID            int64     `gorm:"primaryKey;column:id" json:"-"` // asset id
	WalletAddress string    `gorm:"column:wallet_address" json:"wallet_address"`
	TokenID       int64     `gorm:"column:token_id" json:"token_id"`
	Price         string    `gorm:"column:price" json:"price"`
	Unit          string    `gorm:"column:Unit" json:"unit"`
	Status        uint8     `gorm:"column:status" json:"status"`             // 1:上架 2:下架 3:买 4:卖
	UpdatedTime   time.Time `gorm:"column:updated_time" json:"updated_time"` // update timestamp
	CreatedTime   time.Time `gorm:"column:created_time" json:"created_time"` // create timestamp
}

// TableName get sql table name.获取数据库表名
func (m *TransactionHistory) TableName() string {
	return "transaction_history"
}

// TransactionHistoryColumns get sql column name.获取数据库列名
var TransactionHistoryColumns = struct {
	ID            string
	WalletAddress string
	TokenID       string
	Price         string
	Unit          string
	Status        string
	UpdatedTime   string
	CreatedTime   string
}{
	ID:            "id",
	WalletAddress: "wallet_address",
	TokenID:       "token_id",
	Price:         "price",
	Unit:          "Unit",
	Status:        "status",
	UpdatedTime:   "updated_time",
	CreatedTime:   "created_time",
}
