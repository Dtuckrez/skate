CREATE USER dean WITH LOGIN PASSWORD 'password';
GRANT dean TO postgres;
CREATE DATABASE skate WITH OWNER dean;

\c skate

CREATE TABLE IF NOT EXISTS boards (
    id UUID PRIMARY KEY,
    manufacture VARCHAR(100) NOT NULL,
    created_on TIMESTAMP DEFAULT NOW(),
    updated_on TIMESTAMP NULL);
ALTER TABLE public.boards OWNER to dean;