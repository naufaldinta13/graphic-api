SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `user` (
  `id` BIGINT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(45) NOT NULL,
  `password` VARCHAR(45) NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `last_activity_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `fk_user_1_idx` (`username` ASC));

CREATE TABLE `pasien` (
  `id` BIGINT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `statistic_log` (
  `id` BIGINT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pasien_id` BIGINT(11) UNSIGNED NOT NULL,
  `suhu` VARCHAR(45) NOT NULL,
  `recorded_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_statistic_log_1_idx` (`pasien_id`),
  CONSTRAINT `fk_statistic_log_1` FOREIGN KEY (`pasien_id`) REFERENCES `pasien` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION);


CREATE TABLE `notification` (
  `id` BIGINT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pasien_id` BIGINT(11) UNSIGNED NOT NULL,
  `note` TEXT NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_notification_1_idx` (`pasien_id`),
  CONSTRAINT `fk_notification_1` FOREIGN KEY (`pasien_id`) REFERENCES `pasien` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION);