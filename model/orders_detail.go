package model

import (
	"time"
)

/******sql******
CREATE TABLE `orders_detail` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `order_id` bigint unsigned NOT NULL COMMENT 'orders id',
  `nft_id` bigint NOT NULL,
  `price` varchar(256) NOT NULL,
  `market_type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1:assets 2:avatar',
  `updated_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `created_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`),
  KEY `nft_id` (`nft_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=134 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// OrdersDetail [...]
type OrdersDetail struct {
	ID          uint64    `gorm:"primaryKey;column:id" json:"-"`
	OrderID     uint64    `gorm:"column:order_id" json:"order_id"` // orders id
	NftID       int64     `gorm:"column:nft_id" json:"nft_id"`
	Price       string    `gorm:"column:price" json:"price"`
	MarketType  uint8     `gorm:"column:market_type" json:"market_type"` // 1:assets 2:avatar
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
}

// TableName get sql table name.获取数据库表名
func (m *OrdersDetail) TableName() string {
	return "orders_detail"
}

// OrdersDetailColumns get sql column name.获取数据库列名
var OrdersDetailColumns = struct {
	ID          string
	OrderID     string
	NftID       string
	Price       string
	MarketType  string
	UpdatedTime string
	CreatedTime string
}{
	ID:          "id",
	OrderID:     "order_id",
	NftID:       "nft_id",
	Price:       "price",
	MarketType:  "market_type",
	UpdatedTime: "updated_time",
	CreatedTime: "created_time",
}
