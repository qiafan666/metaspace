ALTER TABLE `metaspace`.`assets`
    ADD COLUMN `sku` varchar(192) NOT NULL AFTER `rarity`;

ALTER TABLE `metaspace`.`assets`
    ADD INDEX `sku`(`sku`) USING BTREE;