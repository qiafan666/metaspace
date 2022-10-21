ALTER TABLE `metaspace`.`assets`
    ADD COLUMN `is_shelf` tinyint(0) UNSIGNED NOT NULL DEFAULT 2 COMMENT '1: is shelf    2:not shelf' AFTER `is_nft`;

UPDATE assets
SET is_shelf = 1
WHERE
        token_id IN (
        SELECT
            orders_detail.nft_id
        FROM
            orders_detail
                INNER JOIN orders ON orders_detail.order_id = orders.id
        WHERE
                orders.`status` = 1
    )