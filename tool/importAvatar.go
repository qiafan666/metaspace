package main

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

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

func main() {

	// go run main.go -db "metaspacebf@metaspace2022@tcp(172.31.36.140:3306)/metaspace?parseTime=true&charset=utf8mb4"
	dbUrl := "metaspacebf@metaspace2022@tcp(172.31.36.140:3306)/metaspace?parseTime=true&charset=utf8mb4"

	// gorm connect mysql
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}

	file, err := os.ReadFile("TotalMetadataSelected.json")
	if err != nil {
		log.Fatalf(err.Error())
	}
	var data []map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var dbIn []Avatar
	for index, v := range data {
		if index < 3 {
			continue
		}
		marshal, err := json.Marshal(v)
		if err != nil {
			log.Fatalf(err.Error())
		}
		dbIn = append(dbIn, Avatar{
			ID:          int64(index),
			Owner:       "none",
			AvatarID:    int64(index),
			Content:     marshal,
			IsMint:      1,
			IsShelf:     2,
			UpdatedTime: time.Now(),
			CreatedTime: time.Now(),
		})
	}
	err = db.Create(&dbIn).Error
	if err != nil {
		log.Fatalf(err.Error())
	}
}
