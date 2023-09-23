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
