CREATE TABLE IF NOT EXISTS account(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    creation_date DATE NOT NULL
);

INSERT INTO account (username, email, password, creation_date) VALUES ('dustinmeyer', 'dustinmeyer10dm@gmail.com', '123', current_time);