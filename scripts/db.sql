CREATE DATABASE IF NOT EXISTS `red_packet_main_db`;

USE `red_packet_main_db`;

CREATE TABLE IF NOT EXISTS `user_tab` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_name` VARCHAR(100) NOT NULL,
    `nick_name` VARCHAR(100) NOT NULL,
    `ctime` INT UNSIGNED NOT NULL,
    `mtime` INT UNSIGNED NOT NULL,
    `extra_data` blob NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `user_wallet_tab` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `balance` BIGINT UNSIGNED NOT NULL,
    `ctime` INT UNSIGNED NOT NULL,
    `mtime` INT UNSIGNED NOT NULL,
    `extra_data` blob NOT NULL,
    INDEX `idx_user_id`(`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `red_packet_tab` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `red_packet_name` VARCHAR(100) NOT NULL DEFAULT '',
    `red_packet_category` INT UNSIGNED NOT NULL,
    `red_packet_result_type` INT UNSIGNED NOT NULL,
    `quantity` INT UNSIGNED NOT NULL,
    `amount` INT UNSIGNED NOT NULL,
    `remaining_quantity` INT UNSIGNED NOT NULL,
    `ctime` INT UNSIGNED NOT NULL,
    `mtime` INT UNSIGNED NOT NULL,
    `extra_data` blob NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `red_packet_transaction_tab` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `transaction_type` INT UNSIGNED NOT NULL,
    `reference_id` VARCHAR(100) NOT NULL DEFAULT '',
    `amount` INT UNSIGNED NOT NULL,
    `ctime` INT UNSIGNED NOT NULL,
    `mtime` INT UNSIGNED NOT NULL,
    `extra_data` blob NOT NULL,
    UNIQUE INDEX `idx_txn_type_ref_id` (`reference_id`, `transaction_type`)  
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;