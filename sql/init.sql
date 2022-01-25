CREATE SCHEMA `efishery_db`;

CREATE TABLE `efishery_db`.`user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(255) NOT NULL,
  `password` VARCHAR (255) NOT NULL,
  `phone` VARCHAR(30) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_username_index` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;