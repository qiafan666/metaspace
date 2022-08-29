ALTER TABLE `metaspace`.`mint_history`
    ADD COLUMN `origin_chain` tinyint(0) UNSIGNED NOT NULL AFTER `token_id`;
ALTER TABLE `metaspace`.`transaction_history`
    ADD COLUMN `origin_chain` tinyint(0) UNSIGNED NOT NULL AFTER `Unit`;