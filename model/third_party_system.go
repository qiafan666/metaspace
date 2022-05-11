package model

import (
	"time"
)

/******sql******
CREATE TABLE `third_party_system` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
  `name` varchar(192) NOT NULL COMMENT 'third party name',
  `apikey` varchar(192) NOT NULL COMMENT 'system create',
  `third_party_public_key` varchar(4096) NOT NULL COMMENT 'rsa public key',
  `callback_address` varchar(192) NOT NULL COMMENT 'call back adress',
  `status` tinyint unsigned NOT NULL COMMENT '1 active 2 inactive',
  `created_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `updated_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/

// ThirdPartySystem [...]
type ThirdPartySystem struct {
	ID                  uint64    `gorm:"primaryKey;column:id" json:"-"`                               // pk
	Name                string    `gorm:"column:name" json:"name"`                                     // third party name
	APIkey              string    `gorm:"column:apikey" json:"apikey"`                                 // system create
	ThirdPartyPublicKey string    `gorm:"column:third_party_public_key" json:"third_party_public_key"` // rsa public key
	CallbackAddress     string    `gorm:"column:callback_address" json:"callback_address"`             // call back adress
	Status              uint8     `gorm:"column:status" json:"status"`                                 // 1 active 2 inactive
	CreatedTime         time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime         time.Time `gorm:"column:updated_time" json:"updated_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ThirdPartySystem) TableName() string {
	return "third_party_system"
}

// ThirdPartySystemColumns get sql column name.获取数据库列名
var ThirdPartySystemColumns = struct {
	ID                  string
	Name                string
	APIkey              string
	ThirdPartyPublicKey string
	CallbackAddress     string
	Status              string
	CreatedTime         string
	UpdatedTime         string
}{
	ID:                  "id",
	Name:                "name",
	APIkey:              "apikey",
	ThirdPartyPublicKey: "third_party_public_key",
	CallbackAddress:     "callback_address",
	Status:              "status",
	CreatedTime:         "created_time",
	UpdatedTime:         "updated_time",
}
