DROP TABLE IF EXISTS cars;
CREATE TABLE cars (
                        id          INT AUTO_INCREMENT NOT NULL,
                        brand        VARCHAR(255) NOT NULL ,
                        model        VARCHAR(255) NOT NULL,
                        country      VARCHAR(255) NOT NULL,
                        year         SMALLINT NOT NULL,
                        status       VARCHAR(255) NOT NULL,
                        enginetype   VARCHAR(255) NOT NULL,
                        enginevolume DECIMAL(65,2) NOT NULL,
                        transmission VARCHAR(255) NOT NULL,
                        drivetype    VARCHAR(255) NOT NULL,
                        color        VARCHAR(255) NOT NULL,
                        milage       DECIMAL(65,2) NOT NULL,
                        other        VARCHAR(1023) NOT NULL,
                        PRIMARY KEY (`id`)
);

INSERT INTO cars
(brand, model, country, year, status, enginetype, enginevolume, transmission, drivetype, color, milage, other)
VALUES
    ('BMW', 'M4','Germeny',2020, 'on the way', 'gasoline',3.0, 'automatic dual clutch', '4wd', 'white', 16342.50, 'available by prepaiment')
