CREATE TABLE `metaspace`.`sign_log`  (
     `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
     `err` varchar(255) NOT NULL,
     `created_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
     `updated_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
     PRIMARY KEY (`id`)
);