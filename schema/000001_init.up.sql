CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    wallet varchar(255) not null default '0x0',
    password_hash varchar(255) not null
);