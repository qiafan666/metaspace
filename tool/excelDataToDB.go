package main

import (
	"flag"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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

	// go run main.go -db "metaspacebf@metaspace2022@tcp(172.31.36.140:3306)/metaspace?parseTime=true&charset=utf8mb4"
	var dbUrl string
	flag.StringVar(&dbUrl, "db", "", "mysql connection url")
	flag.Parse()

	// gorm connect mysql
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		log.Println(err)
		return
	}

	var assets []Assets
	var tokenId, category, subCategory, rarity, indexId, isNft int
	var walletAddress, image, name, description, url, urlContent, originChain string

	fileName := "assets.xlsx"
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowsMap := make(map[int]string)
	rowsMap[1] = "Ticket"
	rowsMap[2] = "Land"
	rowsMap[3] = "Building"
	rowsMap[4] = "Tower"
	rowsMap[5] = "Trap"
	for _, rowName := range rowsMap {

		rows, err := f.GetRows(rowName)
		if err != nil {
			fmt.Println(err)
			return
		}
		var errs error
		for index, row := range rows {
			if index == 0 {
				continue
			}

			tokenId, errs = strconv.Atoi(row[0])
			if errs != nil {
				fmt.Printf("index=%d,row=%d", index, tokenId)
				return
			}

			category, errs = strconv.Atoi(row[1])
			if errs != nil {
				fmt.Printf("index=%d,row=%d", index, category)
				return
			}

			subCategory, errs = strconv.Atoi(row[2])
			if errs != nil {
				fmt.Printf("index=%d,row=%d", index, subCategory)
				return
			}

			rarity, errs = strconv.Atoi(row[3])
			if errs != nil {
				fmt.Printf("index=%d,row=%d", index, rarity)
				return
			}

			indexId, errs = strconv.Atoi(row[14])
			if errs != nil {
				fmt.Printf("index=%d,row=%d", index, indexId)
				return
			}
			isNft, errs = strconv.Atoi(row[4])
			if errs != nil {
				fmt.Printf("index=%d,row=%d", index, indexId)
				return
			}
			if isNft == 0 {
				isNft = 2
			}

			createTime, _ := time.Parse("2006-01-02 15:04:05", row[12])
			updatedTime, _ := time.Parse("2006-01-02 15:04:05", row[13])

			walletAddress = row[5]
			image = row[6]
			name = row[7]
			description = row[8]
			url = row[9]
			urlContent = row[10]
			originChain = row[11]

			assets = append(assets, Assets{
				UID:         walletAddress,
				TokenID:     int64(tokenId),
				Category:    int64(category),
				Type:        int64(subCategory),
				Rarity:      int64(rarity),
				Image:       image,
				Name:        name,
				Description: description,
				URI:         url,
				URIContent:  urlContent,
				OriginChain: originChain,
				IndexID:     uint64(indexId),
				NickName:    name + "#" + row[14],
				IsNft:       uint8(isNft),
				CreatedAt:   createTime,
				UpdatedAt:   updatedTime,
			})
		}

	}

	err = db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&assets).Error
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
