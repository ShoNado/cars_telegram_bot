DROP TABLE IF EXISTS users;
CREATE TABLE users(
    id                  INT AUTO_INCREMENT NOT NULL,
    isAdmin             BOOLEAN DEFAULT false,
    tgId                BIGINT,
    nameFromUser        VARCHAR(255) DEFAULT '',
    nameFromTg          VARCHAR(255) DEFAULT '',
    userName            VARCHAR(255) DEFAULT '',
    phoneNumber         VARCHAR(255) DEFAULT '',
    price               TEXT(10000),
    brandCountryModel   TEXT(10000),
    engine              TEXT(10000),
    transmission        TEXT(10000),
    color               TEXT(10000),
    other               TEXT(20000),
    orderTime           DATETIME,
    IsCompleted         BOOLEAN DEFAULT false,
    IsAdminSaw          BOOLEAN DEFAULT false,
    IsInWork            BOOLEAN DEFAULT false,
    PRIMARY KEY (`id`)
);


