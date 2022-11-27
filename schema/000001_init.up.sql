CREATE TABLE users
(
    id serial not null unique,
    username varchar(255) not null unique,
    wallet varchar(255) not null default '0x0',
    password_hash varchar(255) not null
);


CREATE TABLE transactions
(
    id serial not null unique,
    wallet varchar(255) not null,
    tx_hash varchar(255) not null unique,
    created_at timestamp not null default now(),
    user_id integer not null,
    item_id integer not null
);