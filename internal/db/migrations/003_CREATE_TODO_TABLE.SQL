/* Once time items once they are complete they will be archieved*/

CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(50),
    account_id INTEGER NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account(id)
);