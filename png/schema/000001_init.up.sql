CREATE TABLE `png`
(
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` VARCHAR(50) NOT NULL,
    `image` BLOB NOT NULL
)
;

ALTER TABLE `png`
    ADD INDEX `IXFK_png_user` (`user_id` ASC)
;