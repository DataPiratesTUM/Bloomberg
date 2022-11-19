CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS users (
    id            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    name          VARCHAR (50)    NOT NULL UNIQUE,
    balance       INT             NOT NULL DEFAULT 0 CHECK (balance >= 0)
);

CREATE TABLE IF NOT EXISTS organisations (
    id            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    name          VARCHAR (50)    NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS securities (
    id            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    name          VARCHAR (50)    NOT NULL UNIQUE,
    description   TEXT            NOT NULL,
    creator       uuid            NOT NULL REFERENCES organisations (id),
    creation_date DATE            NOT NULL DEFAULT now(),
    ttl_1         DATE            NOT NULL CHECK (ttl_1 > creation_date),
    ttl_2         DATE            NOT NULL CHECK (ttl_2 > ttl_1)
);

CREATE TABLE IF NOT EXISTS orders (
    id            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    security      uuid            NOT NULL REFERENCES securities (id),
    quantity      INT             NOT NULL CHECK (quantity != 0),
    price         INT             NOT NULL CHECK (price > 0),
    side          BOOLEAN         NOT NULL,
    "user"        uuid            NOT NULL REFERENCES users (id),
    creation_date DATE            NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS matches (
    id            uuid            PRIMARY KEY DEFAULT uuid_generate_v4(),
    buyer         uuid            NOT NULL REFERENCES users (id),
    buy_price     INT             NOT NULL CHECK (buy_price > 0),
    seller        uuid            NOT NULL REFERENCES users (id),
    sell_price    INT             NOT NULL CHECK (sell_price > 0),
    security      uuid            NOT NULL REFERENCES securities (id),
    quantity      INT             NOT NULL CHECK (quantity >= 0)
);

CREATE TABLE IF NOT EXISTS open_orders (
    security      uuid            NOT NULL REFERENCES securities (id),
    quantity      INT             NOT NULL CHECK (quantity >= 0),
    price         INT             NOT NULL CHECK (price > 0),
    side          BOOLEAN         NOT NULL,
    "user"        uuid            NOT NULL REFERENCES users (id),

    PRIMARY KEY(security, price, side, "user")
);


CREATE OR REPLACE FUNCTION order_trigger_function() 
    RETURNS TRIGGER AS $trigger$
BEGIN
    IF NEW.quantity < 0 THEN 
        UPDATE open_orders
        SET quantity = quantity + NEW.quantity
        WHERE security = NEW.security
            AND price = NEW.price
            AND side = NEW.side
            AND "user" = NEW.user;

        IF NOT FOUND THEN 
            RAISE EXCEPTION 'order not found';
        END IF;
    ELSE
        INSERT INTO open_orders (security, quantity, price, side, "user") 
        VALUES (NEW.security, NEW.quantity, NEW.price, NEW.side, NEW.user)
        ON CONFLICT (security, price, side, "user") DO UPDATE 
        SET quantity = open_orders.quantity + EXCLUDED.quantity;
    END IF;

    RETURN NEW;
END; 
$trigger$ LANGUAGE PLPGSQL;

CREATE OR REPLACE TRIGGER order_trigger
    AFTER INSERT
    ON orders
    FOR EACH ROW
    EXECUTE PROCEDURE order_trigger_function();

CREATE OR REPLACE FUNCTION match_trigger_function() 
    RETURNS TRIGGER AS $trigger$
BEGIN
UPDATE open_orders
    SET quantity = quantity - NEW.quantity
    WHERE side
        AND "user" = NEW.buyer
        AND price = NEW.buy_price
        AND security = NEW.security;

    UPDATE open_orders 
    SET quantity = quantity - NEW.quantity
    WHERE NOT side
        AND "user" = NEW.seller
        AND price = NEW.sell_price
        AND security = NEW.security;

    RETURN NEW;
END; 
$trigger$ LANGUAGE PLPGSQL;

CREATE OR REPLACE TRIGGER match_trigger
    AFTER INSERT
    ON matches
    FOR EACH ROW
    EXECUTE PROCEDURE match_trigger_function();
