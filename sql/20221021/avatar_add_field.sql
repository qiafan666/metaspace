ALTER TABLE `metaspace`.`avatar`
    ADD COLUMN `owner` varchar(128) NOT NULL AFTER `id`,
    ADD COLUMN `is_mint` tinyint(0) UNSIGNED NOT NULL COMMENT '1 minted 2 unminted' AFTER `content`,
    ADD COLUMN `is_shelf` tinyint(0) UNSIGNED NOT NULL COMMENT '1:shelf  2:not shelf' AFTER `is_mint`,
    MODIFY COLUMN `avatar_id` bigint(0) NOT NULL AFTER `owner`;
