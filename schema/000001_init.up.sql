CREATE TABLE users
(
    id serial not null unique,
    username varchar(255) not null unique,
    wallet varchar(255) not null default '0x0',
    password_hash varchar(255) not null
);

CREATE TABLE nfts
(
    id serial not null unique,
    nft_id integer not null unique,
    owner varchar(255) not null,
    image varchar(255) not null,
    name varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE items
(
    id serial not null unique,
    item_id integer not null unique,
    nft_id integer not null references nfts(nft_id),
    price decimal not null,
    listing_price decimal not null,
    total_price decimal not null,
    seller_wallet bytea not null,
--     seller_id integer not null references users(id) default 0,
    is_sold bool not null default false
);


CREATE TABLE transactions
(
    id serial not null unique,
    wallet varchar(255) not null,
    tx_hash varchar(255) not null unique,
    created_at timestamp not null default now(),
    user_id integer not null references users(id),
    item_id integer not null references items(item_id),
    nft_id integer not null references nfts(nft_id)
);


SELECT * FROM items it INNER JOIN nfts nft ON it.nft_id = nft.nft_id WHERE it.is_sold=true AND nft.owner='0x70997970C51812dc3A010C7d01b50e0d17dc79C8' ORDER BY it.item_id