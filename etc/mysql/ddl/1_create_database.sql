SET CHARSET UTF8;
CREATE DATABASE IF NOT EXISTS arrogancia DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_bin;
CREATE USER IF NOT EXISTS arrogancia IDENTIFIED BY '5s45adf1hge';
GRANT SELECT, INSERT, UPDATE, DELETE ON mailscat.* TO arrogancia;
