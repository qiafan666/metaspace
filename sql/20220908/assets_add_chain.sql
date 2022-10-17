ALTER TABLE `metaspace`.`assets`
    MODIFY COLUMN `origin_chain` tinyint(16) UNSIGNED NOT NULL COMMENT '1:eth 2:bsc' AFTER `uri_content`;

ALTER TABLE `metaspace`.`assets`
DROP INDEX `token_id`,
ADD UNIQUE INDEX `token_id`(`token_id`) USING BTREE;

ALTER TABLE `metaspace`.`assets`
DROP INDEX `Index_uid`,
DROP INDEX `Index_uid_category_status`;