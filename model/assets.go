package model

import (
	"time"
)

// Assets [...]
type Assets struct {
	Id            int64     `gorm:"primaryKey;column:id" json:"id"`
	Uid           string    `gorm:"column:uid" json:"uid"`
	TokenId       int64     `gorm:"column:token_id" json:"token_id"`
	Category      int64     `gorm:"column:category" json:"category"`
	Type          int64     `gorm:"column:type" json:"type"`
	Rarity        int64     `gorm:"column:rarity" json:"rarity"`
	Image         string    `gorm:"column:image" json:"image"`
	Name          string    `gorm:"column:name" json:"name"`
	Description   string    `gorm:"column:description" json:"description"`
	Uri           string    `gorm:"column:uri" json:"uri"`
	UriContent    string    `gorm:"column:uri_content" json:"uri_content"`
	OriginChain   string    `gorm:"column:origin_chain" json:"origin_chain"`
	BlockNumber   int64     `gorm:"column:block_number" json:"block_number"`
	TxHash        string    `gorm:"column:tx_hash" json:"tx_hash"`
	Status        string    `gorm:"column:status" json:"status"`
	MintSignature string    `gorm:"column:mint_signature" json:"mint_signature"`
	IsNft         uint8     `gorm:"column:is_nft" json:"is_nft"` // 1: is nft    2:not nft
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName get sql table name.
func (m *Assets) TableName() string {
	return "assets"
}

// AssetsColumns get sql column name
var AssetsColumns = struct {
	Id            string
	Uid           string
	TokenId       string
	Category      string
	Type          string
	Rarity        string
	Image         string
	Name          string
	Description   string
	Uri           string
	UriContent    string
	OriginChain   string
	BlockNumber   string
	TxHash        string
	Status        string
	MintSignature string
	IsNft         string
	CreatedAt     string
	UpdatedAt     string
}{
	Id:            "id",
	Uid:           "uid",
	TokenId:       "token_id",
	Category:      "category",
	Type:          "type",
	Rarity:        "rarity",
	Image:         "image",
	Name:          "name",
	Description:   "description",
	Uri:           "uri",
	UriContent:    "uri_content",
	OriginChain:   "origin_chain",
	BlockNumber:   "block_number",
	TxHash:        "tx_hash",
	Status:        "status",
	MintSignature: "mint_signature",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}
