CREATE TABLE `metaspace`.`group`  (
     `id` bigint(0) NOT NULL AUTO_INCREMENT,
     `name` varchar(192) NOT NULL,
     `sku` varchar(192) NOT NULL,
     `created_at` timestamp(3) NOT NULL COMMENT 'create timestamp',
     `updated_at` timestamp(3) NOT NULL COMMENT 'update timestamp',
     PRIMARY KEY (`id`),
     INDEX `index_sku`(`sku`) USING BTREE,
     INDEX `index_group`(`group`) USING BTREE
);