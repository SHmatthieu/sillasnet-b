DROP DATABASE IF EXISTS projecttest;
CREATE DATABASE projecttest;

DROP USER 'usersec1'@'localhost';

CREATE USER 'usersec1'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON projecttest.* TO 'usersec1'@'localhost';

USE projecttest;

DROP TABLE IF EXISTS UserSoftware;
DROP TABLE IF EXISTS UserHardware;
DROP TABLE IF EXISTS UserSupplier;


DROP TABLE IF EXISTS User;
DROP TABLE IF EXISTS Software;
DROP TABLE IF EXISTS Hardware;
DROP TABLE IF EXISTS Supplier;




CREATE TABLE User
(
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR (256) NOT NULL,
    token VARCHAR (256) NOT NULL,
    hashpassword VARCHAR (256) NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (name)
);

CREATE TABLE Software
(
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR (256) NOT NULL,
    version VARCHAR (256) NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (name, version)
);

CREATE TABLE Hardware
(
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR (256) NOT NULL,
    version VARCHAR (256) NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (name, version)
);

CREATE TABLE Supplier
(
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR (256) NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (name)
);


CREATE TABLE UserSoftware
(
    userid INT NOT NULL,
    softwareid INT NOT NULL,
    UNIQUE (userid, softwareid),

    FOREIGN KEY (userid) REFERENCES User(id),
    FOREIGN KEY (softwareid) REFERENCES Software(id)
);

CREATE TABLE UserHardware
(
    userid INT NOT NULL,
    hardwareid INT NOT NULL,
    UNIQUE (userid, hardwareid),

    FOREIGN KEY (userid) REFERENCES User(id),
    FOREIGN KEY (hardwareid) REFERENCES Hardware(id)
);

CREATE TABLE UserSupplier
(
    userid INT NOT NULL,
    supplierid INT NOT NULL,
    UNIQUE (userid, supplierid),

    FOREIGN KEY (userid) REFERENCES User(id),
    FOREIGN KEY (supplierid) REFERENCES Supplier(id)
);


INSERT INTO User (name, token, hashpassword) VALUE ("user1", "token1", "password");
INSERT INTO Software (name, version) VALUE ("soft1", "2.0");
INSERT INTO Software (name, version) VALUE ("soft1", "3.0");
INSERT INTO Software (name, version) VALUE ("soft2", "3.0");
INSERT INTO Hardware (name, version) VALUE ("hard1", "2.0");
INSERT INTO Supplier (name) VALUE ("sup1");

INSERT INTO UserSoftware (userid,softwareid) VALUE (1, 1);
INSERT INTO UserSoftware (userid,softwareid) VALUE (1, 2);

INSERT INTO UserHardware (userid,hardwareid) VALUE (1, 1);
INSERT INTO UserSupplier (userid,supplierid) VALUE (1, 1);