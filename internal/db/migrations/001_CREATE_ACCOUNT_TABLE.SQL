CREATE TABLE IF NOT EXISTS account(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL CHECK (char_length(username) > 5),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    creation_date DATE NOT NULL,
    refresh_token VARCHAR(255) ,
    CONSTRAINT email_format_check CHECK (
        email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'
    )
);

/* INSERT INTO account (username, email, password, creation_date) VALUES ('dustinmeyer', 'dustinmeyer10dm@gmail.com', '123', current_timestamp); */