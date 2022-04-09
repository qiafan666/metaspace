DROP TABLE IF EXISTS `subscribe_newsletter_email`;
CREATE TABLE `subscribe_newsletter_email`  (
                                               `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                                               `email` varchar(192) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                               `status` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '1:可以订阅 2：不可订阅',
                                               `created_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
                                               `updated_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
                                               PRIMARY KEY (`id`) USING BTREE,
                                               UNIQUE INDEX `idx_email_status_id`(`email`, `status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;