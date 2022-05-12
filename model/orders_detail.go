package model

import (
	"time"
)

/******sql******
CREATE TABLE `orders_detail` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `order_id` bigint unsigned NOT NULL COMMENT 'orders id',
  `nft_id` varchar(192) NOT NULL,
  `price` int unsigned NOT NULL,
  `expire_time` timestamp(3) NOT NULL,
  `created_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `updated_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// OrdersDetail [...]
type OrdersDetail struct {
	ID          uint64    `gorm:"primaryKey;column:id" json:"-"`
	OrderID     uint64    `gorm:"column:order_id" json:"order_id"` // orders id
	NftID       string    `gorm:"column:nft_id" json:"nft_id"`
	Price       uint      `gorm:"column:price" json:"price"`
	ExpireTime  time.Time `gorm:"column:expire_time" json:"expire_time"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
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
	ExpireTime  string
	CreatedTime string
	UpdatedTime string
}{
	ID:          "id",
	OrderID:     "order_id",
	NftID:       "nft_id",
	Price:       "price",
	ExpireTime:  "expire_time",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}
