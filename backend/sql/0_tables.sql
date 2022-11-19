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
    seller        uuid            NOT NULL REFERENCES users (id),
    security      uuid            NOT NULL REFERENCES securities (id),
    quantity      INT             NOT NULL CHECK (quantity >= 0),
    price         INT             NOT NULL CHECK (price > 0)
);

CREATE TABLE IF NOT EXISTS open_orders (
    security      uuid            NOT NULL REFERENCES securities (id),
    quantity      INT             NOT NULL CHECK (quantity >= 0),
    price         INT             NOT NULL CHECK (price > 0),
    side          BOOLEAN         NOT NULL,
    "user"        uuid            NOT NULL REFERENCES users (id),

    PRIMARY KEY(security, price, side, "user")
);

CREATE OR REPLACE FUNCTION open_order_trigger_function() 
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

CREATE OR REPLACE TRIGGER open_order_trigger
    AFTER INSERT
    ON orders
    FOR EACH ROW
    EXECUTE PROCEDURE open_order_trigger_function();

CREATE OR REPLACE FUNCTION match_order_trigger_function() 
    RETURNS TRIGGER AS $trigger$
    DECLARE 
        other_price INT;
        other_user uuid;
        amount INT;
        price INT; 
BEGIN
    SELECT 
        o.user, 
        o.price,
        LEAST(o.quantity, NEW.quantity),
        LEAST(o.price, NEW.price) 
    INTO other_user, other_price, amount, price
    FROM open_orders AS o
    WHERE o.side != NEW.side AND o.security = NEW.security
    ORDER BY abs(o.quantity - NEW.quantity) ASC
    LIMIT 1;

    IF NOT FOUND THEN
        RETURN NEW;
    END IF;

    IF NEW.side THEN
        INSERT INTO matches (buyer, seller, security, quantity, price) VALUES (NEW.user, other_user, NEW.security, amount, price);
    ELSE
        INSERT INTO matches (buyer, seller, security, quantity, price) VALUES (other_user, NEW.user, NEW.security, amount, price);
    END IF;

    UPDATE open_orders AS o
    SET quantity = o.quantity - amount    
    WHERE o.security = NEW.security
        AND o.price = other_price
        AND o.side != NEW.side
        AND o.user = other_user;

    NEW.quantity := NEW.quantity - amount;

    RETURN NEW;
END; 
$trigger$ LANGUAGE PLPGSQL;

CREATE OR REPLACE TRIGGER open_order_trigger
    BEFORE INSERT OR UPDATE
    ON open_orders
    FOR EACH ROW
    WHEN (pg_trigger_depth() <= 1)
    EXECUTE PROCEDURE match_order_trigger_function();