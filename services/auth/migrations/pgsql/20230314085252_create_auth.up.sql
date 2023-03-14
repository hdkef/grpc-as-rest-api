CREATE TABLE auth(
    id UUID NOT NULL UNIQUE, 
    user_id UUID NOT NULL UNIQUE, 
    email VARCHAR NOT NULL UNIQUE,
    password TEXT NOT NULL,
    primary key (id)
    );