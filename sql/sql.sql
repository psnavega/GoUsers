CREATE DATABASE IF NOT EXISTS ps-social-media;
USE ps-social-media;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary_key,
    name varchar(50) not null,
    username varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(50) not null unique,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;