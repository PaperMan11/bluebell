CREATE TABLE `user`(
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`user_id` BIGINT(20) NOT NULL,
`username` VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL,
`password` VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL,
`email` VARCHAR(64) COLLATE utf8mb4_general_ci,
`gender` TINYINT(4) NOT NULL DEFAULT '0',
`create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
`update` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
UNIQUE `idx_username` (`username`) USING BTREE,
UNIQUE `idx_user_id` (`user_id`) USING BTREE
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `community`(
`id` INT(11) NOT NULL AUTO_INCREMENT,
`community_id` INT(10) UNSIGNED NOT NULL,
`community_name` VARCHAR(128) COLLATE utf8mb4_general_ci NOT NULL,
`introduction` VARCHAR(256) COLLATE utf8mb4_general_ci NOT NULL,
`create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
`update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY(`id`),
UNIQUE KEY `idx_community_id` (`community_id`),
UNIQUE KEY `idx_community_name` (`community_name`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `community` VALUES ('1', '1', 'Go', 'Golang', '2020-11-11 07:10:10', '2020-11-12 07:10:11');
INSERT INTO `community` VALUES ('2', '2', 'Leetcode', '刷题', '2021-11-11 07:10:10', '2021-11-12 07:10:11');
INSERT INTO `community` VALUES ('3', '3', 'CS:GO', 'Rush B', '2019-11-11 07:10:10', '2019-11-12 07:10:11');
INSERT INTO `community` VALUES ('4', '4', 'LOL', '英雄联盟', '2018-11-11 07:10:10', '2018-11-12 07:10:11');