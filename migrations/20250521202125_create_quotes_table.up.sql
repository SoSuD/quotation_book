CREATE TABLE IF NOT EXISTS quotes (
        id     SERIAL PRIMARY KEY,
        author TEXT   NOT NULL,
        quote  TEXT   NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_quotes_author
        ON quotes(author);
