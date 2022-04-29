package model

import (
	"time"
)

/******sql******
CREATE TABLE `request_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `third_party_id` bigint unsigned NOT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT 'request user id',
  `uri` varchar(192) NOT NULL,
  `parameter` varchar(4096) NOT NULL,
  `created_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `updated_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`),
  KEY `thrid_index` (`third_party_id`,`user_id`,`uri`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// RequestLog [...]
type RequestLog struct {
	ID           uint64    `gorm:"primaryKey;column:id" json:"-"`
	ThirdPartyID uint64    `gorm:"column:third_party_id" json:"third_party_id"`
	UserID       uint64    `gorm:"column:user_id" json:"user_id"` // request user id
	URI          string    `gorm:"column:uri" json:"uri"`
	Parameter    string    `gorm:"column:parameter" json:"parameter"`
	CreatedTime  time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime  time.Time `gorm:"column:updated_time" json:"updated_time"`
}

// TableName get sql table name.获取数据库表名
func (m *RequestLog) TableName() string {
	return "request_log"
}

// RequestLogColumns get sql column name.获取数据库列名
var RequestLogColumns = struct {
	ID           string
	ThirdPartyID string
	UserID       string
	URI          string
	Parameter    string
	CreatedTime  string
	UpdatedTime  string
}{
	ID:           "id",
	ThirdPartyID: "third_party_id",
	UserID:       "user_id",
	URI:          "uri",
	Parameter:    "parameter",
	CreatedTime:  "created_time",
	UpdatedTime:  "updated_time",
}
