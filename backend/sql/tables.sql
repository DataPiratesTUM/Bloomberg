TABLE "security" {
    "id"            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name"          VARCHAR (50)    NOT NULL UNIQUE,
    "description"   TEXT            NOT NULL,
    "creator"       uuid            NOT NULL REFERENCES "organization" ("id"),
    "creation_data" DATE            NOT NULL DEFAULT now(),
    "ttl_1"         DATA            NOT NULL CHECK ("ttl_1" > "creation_data"),
    "ttl_2"         DATA            NOT NULL CHECK ("ttl_2" > "ttl_1")
}

TABLE "user" {
    "id"            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name"          VARCHAR (50)    NOT NULL UNIQUE
}

TABLE "organization" {
    "id"            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name"          VARCHAR (50)    NOT NULL UNIQUE
}

TABLE "orders" {
    "id"            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    "security"      uuid            NOT NULL REFERENCES "security" ("id"),
    "quantity"      INT             NOT NULL CHECK ("quantity" > 0),
    "price"         INT             NOT NULL CHECK ("price" > 0),
    "side"          BOOLEAN         NOT NULL,
    "user"          uuid            NOT NULL REFERENCES "user" ("id")
}

TABLE "matches" {
    "id"            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    "buyer"         uuid            NOT NULL REFERENCES "user" ("id"),
    "seller"        uuid            NOT NULL REFERENCES "user" ("id"),
    "price"         INT             NOT NULL CHECK ("price" > 0),
    "quantity"      INT             NOT NULL CHECK ("quantity" > 0),
    "security"      uuid            NOT NULL REFERENCES "security" ("id")
}