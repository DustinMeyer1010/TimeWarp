/* habits are items that can be complete more than once in a day */

CREATE TABLE IF NOT EXISTS habits_with_time (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL CHECK ("name" <> ''),
    "description" TEXT,
    completion_time INTERVAL NOT NULL CHECK (completion_time >= INTERVAL '10 minutes'),
    account_id INTEGER NOT NULL REFERENCES account(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS habits_time_logs (
    id SERIAL PRIMARY KEY,
    habit_id INT NOT NULL REFERENCES habits_with_time(id) ON DELETE CASCADE,
    "current_time" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    time_spent INTERVAL NOT NULL DEFAULT '0 hours 0 minutes'
);

CREATE TABLE IF NOT EXISTS habits_with_time_completed (
    id SERIAL PRIMARY KEY,
    habit_id INT NOT NULL REFERENCES habits_with_time(id) ON DELETE CASCADE,
    completion_date DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE IF NOT EXISTS habits_without_time (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL CHECK ("name" <> ''),
    "description" TEXT,
    account_id INTEGER NOT NULL REFERENCES account(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS habits_without_time_completed (
    id SERIAL PRIMARY KEY,
    habit_id INT NOT NULL REFERENCES habits_without_time(id) ON DELETE CASCADE,
    completion_date DATE NOT NULL DEFAULT CURRENT_DATE
);
