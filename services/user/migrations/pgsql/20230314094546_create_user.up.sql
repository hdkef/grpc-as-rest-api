CREATE TABLE users(
    id UUID NOT NULL UNIQUE, 
    email VARCHAR NOT NULL UNIQUE,
    name VARCHAR NOT NULL,
    address TEXT NOT NULL,
    primary key (id)
);