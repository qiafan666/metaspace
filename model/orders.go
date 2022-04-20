package model

import (
	"time"
)

/******sql******
CREATE TABLE `orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `seller` varchar(192) NOT NULL,
  `buyer` varchar(192) DEFAULT NULL,
  `signature` varchar(192) NOT NULL,
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1:active 2:expire 3:canceled 4:finished',
  `created_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `updated_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Orders [...]
type Orders struct {
	ID          uint64    `gorm:"primaryKey;column:id" json:"-"`
	Seller      string    `gorm:"column:seller" json:"seller"`
	Buyer       string    `gorm:"column:buyer" json:"buyer"`
	Signature   string    `gorm:"column:signature" json:"signature"`
	Status      uint8     `gorm:"column:status" json:"status"` // 1:active 2:expire 3:canceled 4:finished
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
}

// TableName get sql table name.获取数据库表名
func (m *Orders) TableName() string {
	return "orders"
}

// OrdersColumns get sql column name.获取数据库列名
var OrdersColumns = struct {
	ID          string
	Seller      string
	Buyer       string
	Signature   string
	Status      string
	CreatedTime string
	UpdatedTime string
}{
	ID:          "id",
	Seller:      "seller",
	Buyer:       "buyer",
	Signature:   "signature",
	Status:      "status",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}
