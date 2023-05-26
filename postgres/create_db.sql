CREATE TABLE IF NOT EXISTS "User" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    family VARCHAR(128) NOT NULL,
    age INTEGER CHECK (age > 0),
    sex VARCHAR(10) CHECK (sex IN ('male', 'female')),
    createdAt TIMESTAMPTZ DEFAULT NOW()
);