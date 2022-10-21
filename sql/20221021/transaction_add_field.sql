ALTER TABLE `metaspace`.`transaction_history`
    ADD COLUMN `market_type` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '1:assets 2:avatar' AFTER `status`;