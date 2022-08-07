CREATE TABLE IF NOT EXISTS `user_tab` (
    `user_id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_name` VARCHAR(20) NOT NULL,
    `nick_name` VARCHAR(100) NOT NULL,
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `red_packet_tab` (
    `red_packet_id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `red_packet_name` VARCHAR(20) NOT NULL DEFAULT '',
    `red_packet_type` INT UNSIGNED NOT NULL,
    `amount` INT UNSIGNED NOT NULL,
    `count` INT UNSIGNED NOT NULL, 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;