CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "users" (
    "id"            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name"          VARCHAR (50)    NOT NULL UNIQUE,
    "balance"       INT             NOT NULL DEFAULT 0 CHECK ("balance" >= 0)
);

CREATE TABLE IF NOT EXISTS "organisation" (
    "id"            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name"          VARCHAR (50)    NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "security" (
    "id"            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name"          VARCHAR (50)    NOT NULL UNIQUE,
    "description"   TEXT            NOT NULL,
    "creator"       uuid            NOT NULL REFERENCES "organisation" ("id"),
    "creation_date" DATE            NOT NULL DEFAULT now(),
    "ttl_1"         DATE            NOT NULL CHECK ("ttl_1" > "creation_date"),
    "ttl_2"         DATE            NOT NULL CHECK ("ttl_2" > "ttl_1")
);

CREATE TABLE IF NOT EXISTS "order" (
    "id"            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    "type"          BOOLEAN         NOT NULL,
    "security"      uuid            NOT NULL REFERENCES "security" ("id"),
    "quantity"      INT             NOT NULL CHECK ("quantity" > 0),
    "price"         INT             NOT NULL CHECK ("price" > 0),
    "side"          BOOLEAN         NOT NULL,
    "users"          uuid            NOT NULL REFERENCES "users" ("id"),
    "creation_date" DATE            NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "matches" (
    "id"            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    "buyer"         uuid            NOT NULL REFERENCES "order" ("id"),
    "seller"        uuid            NOT NULL REFERENCES "order" ("id")
);
