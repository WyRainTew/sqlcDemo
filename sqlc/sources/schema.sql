CREATE TYPE book_category AS ENUM (
    'fiction',
    'non_fiction',
    'science',
    'technology',
    'art',
    'history',
    'biography',
    'children',
    'education',
    'reference'
);

CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       description TEXT NOT NULL,
                       metadata BYTEA,
                       category book_category NOT NULL,
                       price DECIMAL(10, 2) NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP
);