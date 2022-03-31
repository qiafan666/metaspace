package model

import (
	"time"
)

type SubscribeNewsletterEmail struct {
	ID          uint64    `gorm:"primaryKey;column:id" json:"-"`
	Email       string    `gorm:"column:email" json:"email"`   // email
	Status      uint8     `gorm:"column:status" json:"status"` // 1：已订阅  2：取消订阅
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
}

// TableName get sql table name.获取数据库表名
func (m *SubscribeNewsletterEmail) TableName() string {
	return "subscribe_newsletter_email"
}

// UserColumns get sql column name.获取数据库列名
var SubscribeNewsletterEmailColumns = struct {
	ID          string
	Email       string
	Status      string
	CreatedTime string
	UpdatedTime string
}{
	ID:          "id",
	Email:       "email",
	Status:      "status",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
}
