package model

import (
	"time"
)

/******sql******
CREATE TABLE `assets` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'asset id',
  `uid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'user id',
  `uuid` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT 'third_party association',
  `token_id` bigint NOT NULL COMMENT 'token id of erc721; should be the same as id',
  `category` bigint NOT NULL COMMENT 'category',
  `type` bigint NOT NULL COMMENT 'type',
  `rarity` bigint NOT NULL DEFAULT '0' COMMENT 'rarity',
  `image` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT 'image',
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT 'name',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT 'description',
  `uri` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT 'uri',
  `uri_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT 'uri content',
  `origin_chain` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'origin chain',
  `block_number` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT 'block number',
  `tx_hash` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT 'transaction hash',
  `status` tinyint unsigned DEFAULT NULL COMMENT 'status',
  `mint_signature` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT 'mint signature',
  `is_nft` tinyint unsigned NOT NULL DEFAULT '2' COMMENT '1: is nft    2:not nft',
  `created_at` datetime(3) NOT NULL COMMENT 'create timestamp',
  `updated_at` datetime(3) NOT NULL COMMENT 'update timestamp',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `Index_uid` (`uid`) USING BTREE,
  KEY `Index_uid_category_status` (`uid`,`category`,`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=323 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='asset table'
******sql******/
// Assets asset table
type Assets struct {
	ID            int64     `gorm:"primaryKey;column:id" json:"-"`               // asset id
	UID           string    `gorm:"column:uid" json:"uid"`                       // user id
	UUID          string    `gorm:"column:uuid" json:"uuid"`                     // third_party association
	TokenID       int64     `gorm:"column:token_id" json:"token_id"`             // token id of erc721; should be the same as id
	Category      int64     `gorm:"column:category" json:"category"`             // category
	Type          int64     `gorm:"column:type" json:"type"`                     // type
	Rarity        int64     `gorm:"column:rarity" json:"rarity"`                 // rarity
	Image         string    `gorm:"column:image" json:"image"`                   // image
	Name          string    `gorm:"column:name" json:"name"`                     // name
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

// AssetsColumns get sql column name.获取数据库列名
var AssetsColumns = struct {
	ID            string
	UID           string
	UUID          string
	TokenID       string
	Category      string
	Type          string
	Rarity        string
	Image         string
	Name          string
	Description   string
	URI           string
	URIContent    string
	OriginChain   string
	BlockNumber   string
	TxHash        string
	Status        string
	MintSignature string
	IsNft         string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	UID:           "uid",
	UUID:          "uuid",
	TokenID:       "token_id",
	Category:      "category",
	Type:          "type",
	Rarity:        "rarity",
	Image:         "image",
	Name:          "name",
	Description:   "description",
	URI:           "uri",
	URIContent:    "uri_content",
	OriginChain:   "origin_chain",
	BlockNumber:   "block_number",
	TxHash:        "tx_hash",
	Status:        "status",
	MintSignature: "mint_signature",
	IsNft:         "is_nft",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}
