CREATE TABLE IF NOT EXISTS habit(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    account_id INTEGER,
    FOREIGN KEY (account_id) REFERENCES account(id)
);