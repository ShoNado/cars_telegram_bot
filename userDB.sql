DROP TABLE IF EXISTS users;
CREATE TABLE users(
    id                  INT AUTO_INCREMENT NOT NULL,
    isAdmin             BOOLEAN,
    tgId                BIGINT,
    nameFromUser        VARCHAR(255),
    nameFromTg          VARCHAR(255),
    userName            VARCHAR(255),
    phoneNumber         VARCHAR(255),
    price               TEXT(10000),
    brandCountryModel   TEXT(10000),
    engine              TEXT(10000),
    transmission        TEXT(10000),
    color               TEXT(10000),
    other               TEXT(20000),
    orderTime           DATETIME,
    IsCompleted         BOOLEAN,
    IsAdminSaw          BOOLEAN,
    IsInWork            BOOLEAN,
    PRIMARY KEY (`id`)
);

INSERT INTO users
(isAdmin,tgId,nameFromUser,userName)
VALUES
#(true,231043417,'Евгений','ShoNado69'),
(true,314539937,'Дмитрий','blyaD1ma')

