package model

import (
	"time"
)

/******sql******
CREATE TABLE `avatar` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `owner` varchar(128) NOT NULL,
  `avatar_id` bigint NOT NULL,
  `content` json NOT NULL,
  `is_mint` tinyint unsigned NOT NULL COMMENT '1 minted 2 unminted',
  `is_shelf` tinyint unsigned NOT NULL COMMENT '1:shelf  2:not shelf',
  `updated_time` timestamp(3) NOT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'update timestamp',
  `created_time` timestamp(3) NOT NULL COMMENT 'create timestamp',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_avatar` (`avatar_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Avatar [...]
type Avatar struct {
	ID          int64     `gorm:"primaryKey;column:id" json:"-"` // id
	Owner       string    `gorm:"column:owner" json:"owner"`
	AvatarID    int64     `gorm:"column:avatar_id" json:"avatar_id"`
	Content     []byte    `gorm:"column:content" json:"content"`
	IsMint      uint8     `gorm:"column:is_mint" json:"is_mint"`           // 1 minted 2 unminted
	IsShelf     uint8     `gorm:"column:is_shelf" json:"is_shelf"`         // 1:shelf  2:not shelf
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"` // update timestamp
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"` // create timestamp
}

// TableName get sql table name.获取数据库表名
func (m *Avatar) TableName() string {
	return "avatar"
}

// AvatarColumns get sql column name.获取数据库列名
var AvatarColumns = struct {
	ID          string
	Owner       string
	AvatarID    string
	Content     string
	IsMint      string
	IsShelf     string
	UpdatedTime string
	CreatedTime string
}{
	ID:          "id",
	Owner:       "owner",
	AvatarID:    "avatar_id",
	Content:     "content",
	IsMint:      "is_mint",
	IsShelf:     "is_shelf",
	UpdatedTime: "updated_time",
	CreatedTime: "created_time",
}
