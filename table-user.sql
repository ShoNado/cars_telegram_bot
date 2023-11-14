DROP TABLE IF EXISTS users;
CREATE TABLE users(
    id              INT AUTO_INCREMENT NOT NULL,
    isAdmin         BOOLEAN NOT NULL,
    tgId            INT NOT NULL,
    firstName       VARCHAR(255),
    secondName      VARCHAR(255),
    userName        VARCHAR(255),
    modelBrand      VARCHAR(2047),
    country         VARCHAR(255),
    enigine         VARCHAR(2047),
    transmission    VARCHAR(255),
    drivetype       VARCHAR(255),
    color           VARCHAR(255),
    price           varchar(2047),
    other           VARCHAR(1023),
    favorite
    IsCompleted     BOOLEAN,
    PRIMARY KEY (`id`)
);