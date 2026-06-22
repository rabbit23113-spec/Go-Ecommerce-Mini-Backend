CREATE TABLE
    users (
        id SERIAL PRIMARY KEY DEFAULT UUID (),
        role VARCHAR DEFAULT 'customer',
        email VARCHAR UNIQUE NOT NULL,
        password_hash VARCHAR NOT NULL,
        created_at TIMESTAMPTZ NOW ()
    )
CREATE TABLE
    auth_sessions (
        id SERIAL PRIMARY KEY DEFAULT UUID (),
        user_id VARCHAR NOT NULL,
        token_hash VARCHAR NOT NULL,
        revoked_at TIMESTAMPTZ DEFAULT NULL,
        created_at TIMESTAMPTZ NOW ()
    )
CREATE TABLE
    products (
        id SERIAL PRIMARY KEY DEFAULT UUID (),
        name VARCHAR NOT NULL,
        price INT NOT NULL,
        created_at TIMESTAMPTZ NOW ()
    )
CREATE TABLE
    warehouses (
        id SERIAL PRIMARY KEY DEFAULT UUID (),
        location VARCHAR NOT NULL,
        status VARCHAR NOT NULL,
        product_ids VARCHAR DEFAULT '[]',
        created_at TIMESTAMPTZ NOW ()
    )
CREATE TABLE
    orders (
        id SERIAL PRIMARY KEY DEFAULT UUID (),
        user_id VARCHAR NOT NULL,
        product_id VARCHAR NOT NULL,
        quantity INT NOT NULL,
        status VARCHAR DEFAULT 'in_process',
        created_at TIMESTAMPTZ NOW ()
    )