CREATE TABLE `user`
(
    `id` VARCHAR(50) NOT NULL,
    `login` VARCHAR(50) NOT NULL UNIQUE ,
    `password` VARCHAR(50) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY (id)
)
;