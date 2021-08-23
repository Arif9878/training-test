CREATE DATABASE  IF NOT EXISTS `test`;
USE `test`;

DROP TABLE IF EXISTS `log_search`;

CREATE TABLE `log_search` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    datetime DATETIME,
    search CHAR(255)
);