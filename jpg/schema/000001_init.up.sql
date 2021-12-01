CREATE TABLE `jpg`
(
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` VARCHAR(50) NOT NULL,
    `image` BLOB NOT NULL
);

ALTER TABLE `jpg`
    ADD INDEX `IXFK_jpg_user` (`user_id` ASC)
;