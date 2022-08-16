package model

import (
	"time"
)

/******sql******
CREATE TABLE `avatar` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `avatar_id` bigint unsigned NOT NULL,
  `content` json NOT NULL,
  `updated_time` timestamp(3) NOT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'update timestamp',
  `created_time` timestamp(3) NOT NULL COMMENT 'create timestamp',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_avatar` (`avatar_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Avatar [...]
type Avatar struct {
	ID          int64     `gorm:"primaryKey;column:id" json:"-"` // id
	AvatarID    uint64    `gorm:"column:avatar_id" json:"avatar_id"`
	Content     []byte    `gorm:"column:content" json:"content"`
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
	AvatarID    string
	Content     string
	UpdatedTime string
	CreatedTime string
}{
	ID:          "id",
	AvatarID:    "avatar_id",
	Content:     "content",
	UpdatedTime: "updated_time",
	CreatedTime: "created_time",
}
