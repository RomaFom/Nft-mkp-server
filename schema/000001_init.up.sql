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
    nft_id int not null unique,
    owner_id varchar(255) not null,
    image varchar(255) not null,
    name varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE items
(
    id serial not null unique,
    item_id varchar(255) not null unique,
    nft_id int not null references nfts(nft_id),
    price decimal not null,
    listing_price decimal not null,
    total_price decimal not null,
    seller_wallet varchar(255) not null,
    seller_id int not null references users(id),
    is_sold boolean not null default false
);


CREATE TABLE transactions
(
    id serial not null unique,
    wallet varchar(255) not null,
    tx_hash varchar(255) not null unique,
    created_at timestamp not null default now(),
    user_id integer not null references users(id),
    item_id integer not null references items(id),
    nft_id integer not null references nfts(nft_id)
);