package model

import (
	"time"
)

/******sql******
CREATE TABLE `orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `seller` varchar(192) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_as_ci NOT NULL,
  `buyer` varchar(192) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `signature` varchar(192) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '经过metamask的上架签名',
  `salt_nonce` varchar(192) NOT NULL COMMENT '上架签名中的salt_nonce',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1:active 2:expire 3:canceled 4:finished',
  `total_price` varchar(256) NOT NULL,
  `start_time` timestamp(3) NOT NULL,
  `expire_time` timestamp(3) NOT NULL,
  `updated_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `created_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Orders [...]
type Orders struct {
	ID          uint64    `gorm:"primaryKey;column:id" json:"-"`
	Seller      string    `gorm:"column:seller" json:"seller"`
	Buyer       string    `gorm:"column:buyer" json:"buyer"`
	Signature   string    `gorm:"column:signature" json:"signature"`   // 经过metamask的上架签名
	SaltNonce   string    `gorm:"column:salt_nonce" json:"salt_nonce"` // 上架签名中的salt_nonce
	Status      uint8     `gorm:"column:status" json:"status"`         // 1:active 2:expire 3:canceled 4:finished
	TotalPrice  string    `gorm:"column:total_price" json:"total_price"`
	StartTime   time.Time `gorm:"column:start_time" json:"start_time"`
	ExpireTime  time.Time `gorm:"column:expire_time" json:"expire_time"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
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
	SaltNonce   string
	Status      string
	TotalPrice  string
	StartTime   string
	ExpireTime  string
	UpdatedTime string
	CreatedTime string
}{
	ID:          "id",
	Seller:      "seller",
	Buyer:       "buyer",
	Signature:   "signature",
	SaltNonce:   "salt_nonce",
	Status:      "status",
	TotalPrice:  "total_price",
	StartTime:   "start_time",
	ExpireTime:  "expire_time",
	UpdatedTime: "updated_time",
	CreatedTime: "created_time",
}
