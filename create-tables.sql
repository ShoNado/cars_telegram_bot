DROP TABLE IF EXISTS cars;
CREATE TABLE cars (
                        id          INT AUTO_INCREMENT NOT NULL,
                        brand        VARCHAR(255) NOT NULL,
                        model        VARCHAR(255) NOT NULL,
                        country      VARCHAR(255),
                        year         SMALLINT,
                        status       VARCHAR(255),
                        enginetype   VARCHAR(255),
                        enginevolume VARCHAR(255),
                        horsepower   VARCHAR(255),
                        torque       VARCHAR(255),
                        transmission VARCHAR(255),
                        drivetype    VARCHAR(255),
                        color        VARCHAR(255),
                        milage       VARCHAR(255),
                        price        varchar(255),
                        other        VARCHAR(1023),
                        IsCompleted  BOOLEAN,
                        PRIMARY KEY (`id`)
);

INSERT INTO cars
(brand, model, country, year, status, enginetype, enginevolume,horsepower, transmission,torque, drivetype, color, milage, price, other, IsCompleted)
VALUES
    ('BMW', 'M4','Germeny',2020, 'on the way', 'gasoline','3.0', '525','600','automatic dual clutch', '4wd', 'white', '16342.50', '120.000$', 'available by prepaiment', 1),
    ('Mercedes', 'E220e','Germeny',2018, 'available', 'hybrid', '2.0','199','240', 'automatic dual clutch', '4wd', 'black', '207342.50', '40.000$', 'available by prepaiment', 1)