
ALTER TABLE `metaspace`.`assets` 
MODIFY COLUMN `origin_chain` bigint(0) UNSIGNED NOT NULL AFTER `uri_content`;
ALTER TABLE `metaspace`.`transaction_history` 
MODIFY COLUMN `origin_chain` bigint(0) UNSIGNED NOT NULL AFTER `Unit`;
ALTER TABLE `metaspace`.`mint_history` 
MODIFY COLUMN `origin_chain` bigint(0) UNSIGNED NOT NULL AFTER `token_id`;
