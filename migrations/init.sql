CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    description TEXT,
    banner_url TEXT,
    category TEXT,
    author TEXT,
    datetime TIMESTAMPTZ,
    address TEXT
    location TEXT,
    city TEXT,
);

CREATE TABLE promotion (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255),
    banner_url TEXT,
    reviews_number INTEGER,
    reviews_rate NUMERIC(3, 2),
    expires TIMESTAMP WITH TIME ZONE,
    discount INTEGER,
    city VARCHAR(255),
    address VARCHAR(255),
    latitude NUMERIC(10, 6),
    longitude NUMERIC(10, 6),
    price NUMERIC(10, 2)
);
