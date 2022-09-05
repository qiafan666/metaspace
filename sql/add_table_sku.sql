CREATE TABLE `metaspace`.`sku`  (
    `id` bigint(0) NOT NULL AUTO_INCREMENT,
    `name` varchar(192) NOT NULL,
    `category` bigint(0) NOT NULL COMMENT 'category',
    `type` bigint(0) NOT NULL COMMENT 'type',
    `rarity` bigint(0) NOT NULL COMMENT 'rarity',
    `created_at` timestamp(3) NOT NULL COMMENT 'create timestamp',
    `updated_at` timestamp(3) NOT NULL COMMENT 'update timestamp',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `index_unique`(`category`, `type`, `rarity`) USING BTREE,
    UNIQUE INDEX `index_sku`(`sku`) USING BTREE
);
