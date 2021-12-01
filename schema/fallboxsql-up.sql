SET FOREIGN_KEY_CHECKS=0
; 
/* Drop Tables */

DROP TABLE IF EXISTS `jpg` CASCADE
;

DROP TABLE IF EXISTS `png` CASCADE
;

DROP TABLE IF EXISTS `user` CASCADE
;

/* Create Tables */

CREATE TABLE `jpg`
(
	`id` INT NOT NULL AUTO_INCREMENT,
	`user_id` VARCHAR(50) NOT NULL,
	`image` BLOB NOT NULL,
	CONSTRAINT `PK_jpg` PRIMARY KEY (`id` ASC)
)

;

CREATE TABLE `png`
(
	`id` INT NOT NULL AUTO_INCREMENT,
	`user_id` VARCHAR(50) NOT NULL,
	`image` BLOB NOT NULL,
	CONSTRAINT `PK_png` PRIMARY KEY (`id` ASC)
)

;

CREATE TABLE `user`
(
	`id` VARCHAR(50) NOT NULL,
	`login` VARCHAR(50) NOT NULL,
	`password` VARCHAR(50) NOT NULL,
	CONSTRAINT `PK_Table A` PRIMARY KEY (`id` ASC)
)

;

/* Create Primary Keys, Indexes, Uniques, Checks */

ALTER TABLE `jpg` 
 ADD INDEX `IXFK_jpg_user` (`user_id` ASC)
;

ALTER TABLE `png` 
 ADD INDEX `IXFK_png_user` (`user_id` ASC)
;

/* Create Foreign Key Constraints */

ALTER TABLE `jpg` 
 ADD CONSTRAINT `FK_jpg_user`
	FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
;

ALTER TABLE `png` 
 ADD CONSTRAINT `FK_png_user`
	FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
;

SET FOREIGN_KEY_CHECKS=1
; 
