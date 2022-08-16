CREATE TABLE `metaspace`.`avatar`  (
   `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'id' FIRST,
   `avatar_id` bigint(0) UNSIGNED NOT NULL AFTER `id`,
   `message` json NOT NULL AFTER `avatar_id`,
   `updated_time` timestamp(3) NOT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'update timestamp' AFTER `message`,
   `created_time` timestamp(3) NOT NULL COMMENT 'create timestamp' AFTER `updated_time`,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `index_avatar`(`avatar_id`) USING BTREE
);