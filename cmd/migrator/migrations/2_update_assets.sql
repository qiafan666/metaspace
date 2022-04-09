ALTER TABLE `metaspace`.`assets`
    MODIFY COLUMN `status` tinyint(0) UNSIGNED NULL DEFAULT NULL COMMENT 'status' AFTER `tx_hash`;