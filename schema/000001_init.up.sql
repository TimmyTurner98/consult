CREATE TABLE USERS
(
    id SERIAL PRIMARY KEY,
    name varchar(255) not null,
    email varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE CLIENTS
(
    id          SERIAL PRIMARY KEY,
    name        varchar(255) ,
    number      varchar(255)  not null unique,
    companyname varchar(255) ,
    companybin  varchar(255) unique ,
    client_type  varchar(255)
);
