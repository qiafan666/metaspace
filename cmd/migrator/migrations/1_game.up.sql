BEGIN;

-- ----------------------------
-- Table structure for assets
-- ----------------------------
DROP TABLE IF EXISTS `assets`;
CREATE TABLE `assets` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'asset id',
  `uid` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'user id',
  `token_id` bigint NOT NULL COMMENT 'token id of erc721; should be the same as id',
  `category` int(8) NOT NULL COMMENT 'category',
  `type` int(8) NOT NULL COMMENT 'type',
  `rarity` int(8) NOT NULL DEFAULT 0 COMMENT 'rarity',
  `image` text COLLATE utf8mb4_bin COMMENT 'image',
  `name` text COLLATE utf8mb4_bin COMMENT 'name',
  `description` text COLLATE utf8mb4_bin COMMENT 'description',
  `uri` text COLLATE utf8mb4_bin COMMENT 'uri',
  `uri_content` text COLLATE utf8mb4_bin COMMENT 'uri content',
  `origin_chain` varchar(16) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'origin chain',
  `block_number` text COLLATE utf8mb4_bin COMMENT 'block number',
  `tx_hash` text COLLATE utf8mb4_bin COMMENT 'transaction hash',
  `status` varchar(64) COLLATE utf8mb4_bin COMMENT 'status',
  `created_at` datetime(3) NOT NULL COMMENT 'create timestamp',
  `updated_at` datetime(3) NOT NULL COMMENT 'update timestamp',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `Index_uid` (`uid`) USING BTREE,
  KEY `Index_uid_category_status` (`uid`, `category`, `status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='asset table';