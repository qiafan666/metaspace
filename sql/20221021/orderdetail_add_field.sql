ALTER TABLE `metaspace`.`orders_detail`
    ADD COLUMN `market_type` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '1:assets 2:avatar' AFTER `price`;
