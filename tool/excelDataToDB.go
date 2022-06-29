package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xuri/excelize/v2"
	"strconv"
	"time"
)

type Assets struct {
	ID            int64     `gorm:"primaryKey;column:id" json:"-"`   // asset id
	UID           string    `gorm:"column:uid" json:"uid"`           // user id
	UUID          string    `gorm:"column:uuid" json:"uuid"`         // third_party association
	TokenID       int64     `gorm:"column:token_id" json:"token_id"` // token id of erc721; should be the same as id
	Category      int64     `gorm:"column:category" json:"category"` // category
	Type          int64     `gorm:"column:type" json:"type"`         // type
	Rarity        int64     `gorm:"column:rarity" json:"rarity"`     // rarity
	Image         string    `gorm:"column:image" json:"image"`       // image
	Name          string    `gorm:"column:name" json:"name"`         // name
	IndexID       uint64    `gorm:"column:index_id" json:"index_id"`
	NickName      string    `gorm:"column:nick_name" json:"nick_name"`           // nick name
	Description   string    `gorm:"column:description" json:"description"`       // description
	URI           string    `gorm:"column:uri" json:"uri"`                       // uri
	URIContent    string    `gorm:"column:uri_content" json:"uri_content"`       // uri content
	OriginChain   string    `gorm:"column:origin_chain" json:"origin_chain"`     // origin chain
	BlockNumber   string    `gorm:"column:block_number" json:"block_number"`     // block number
	TxHash        string    `gorm:"column:tx_hash" json:"tx_hash"`               // transaction hash
	Status        uint8     `gorm:"column:status" json:"status"`                 // status
	MintSignature string    `gorm:"column:mint_signature" json:"mint_signature"` // mint signature
	IsNft         uint8     `gorm:"column:is_nft" json:"is_nft"`                 // 1: is nft    2:not nft
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`         // create timestamp
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`         // update timestamp
}

// TableName get sql table name.获取数据库表名
func (m *Assets) TableName() string {
	return "assets"
}

//excel data insert to db
func main() {

	db, err := gorm.Open("mysql", "root:!devpass123456@tcp(3.20.122.137:3306)/metaspacetest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}

	var assets []Assets

	for i := 2; i < 7; i++ {
		fileName := "assets" + strconv.Itoa(i) + ".xlsx"
		f, err := excelize.OpenFile(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		rows, err := f.GetRows("assets")
		if err != nil {
			fmt.Println(err)
			return
		}
		for index, row := range rows {
			if index == 0 {
				continue
			}

			row3, _ := strconv.Atoi(row[2])
			row4, _ := strconv.Atoi(row[3])
			row5, _ := strconv.Atoi(row[4])
			row6, _ := strconv.Atoi(row[5])
			row22, _ := strconv.Atoi(row[21])
			assets = append(assets, Assets{
				UID:         row[1],
				TokenID:     int64(row3),
				Category:    int64(row4),
				Type:        int64(row5),
				Rarity:      int64(row6),
				Image:       row[7],
				Name:        row[8],
				Description: row[9],
				URI:         row[10],
				URIContent:  row[11],
				OriginChain: row[12],
				IndexID:     uint64(row22),
				NickName:    row[8] + "#" + row[21],
				IsNft:       2,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			})
		}

	}

	err = db.Transaction(func(tx *gorm.DB) error {
		for k, _ := range assets {
			err2 := tx.Create(&assets[k]).Error
			if err2 != nil {
				return err2
			}
			fmt.Println(k)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
