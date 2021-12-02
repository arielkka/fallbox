CREATE TABLE `excel`
(
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` VARCHAR(50) NOT NULL,
    `body` LONGBLOB NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY (id)
)
;