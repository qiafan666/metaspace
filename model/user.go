package model

import (
	"time"
)

/******sql******
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `UUID` varchar(192) NOT NULL,
  `email` varchar(192) NOT NULL COMMENT 'email',
  `wallet_address` varchar(192) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'wallet address',
  `password` varchar(192) NOT NULL COMMENT 'password',
  `created_at` timestamp(3) NOT NULL,
  `updated_at` timestamp(3) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/

// User [...]
type User struct {
	ID            uint64    `gorm:"primaryKey;column:id" json:"-"`
	UUID          string    `gorm:"column:UUID" json:"uuid"`
	Email         string    `gorm:"column:email" json:"email"`                   // email
	WalletAddress string    `gorm:"column:wallet_address" json:"wallet_address"` // wallet address
	Password      string    `gorm:"column:password" json:"password"`             // password
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID            string
	UUID          string
	Email         string
	WalletAddress string
	Password      string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	UUID:          "UUID",
	Email:         "email",
	WalletAddress: "wallet_address",
	Password:      "password",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}
