DROP DATABASE IF EXISTS projecttest2;
CREATE DATABASE projecttest2;

DROP USER 'usersec1'@'localhost';

CREATE USER 'usersec1'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON projecttest2.* TO 'usersec1'@'localhost';

USE projecttest2;
