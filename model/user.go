package model

import (
	"time"
)

/******sql******
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `UUID` varchar(192) NOT NULL,
  `email` varchar(192) NOT NULL COMMENT 'email',
  `password` varchar(192) NOT NULL COMMENT 'password',
  `created_at` timestamp(3) NOT NULL,
  `updated_at` timestamp(3) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/

// User [...]
type User struct {
	ID            uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"-"`
	UUID          string    `gorm:"column:UUID;type:varchar(192);not null" json:"uuid"`
	Email         string    `gorm:"column:email;type:varchar(192);not null" json:"email"`                   // email
	WalletAddress string    `gorm:"column:wallet_address;type:varchar(192);not null" json:"wallet_address"` // email
	Password      string    `gorm:"column:password;type:varchar(192);not null" json:"password"`             // password
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp(3);not null" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp(3);not null" json:"updatedAt"`
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
