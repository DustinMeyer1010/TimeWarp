/* habits are items that can be complete more than once in a day */

CREATE TABLE IF NOT EXISTS habits (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    completion_time INTERVAL NOT NULL,
    account_id INTEGER NOT NULL REFERENCES account(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS habits_time_logs (
    id SERIAL PRIMARY KEY,
    habit_id INT NOT NULL REFERENCES habits(id) ON DELETE CASCADE,
    day DATE NOT NULL DEFAULT CURRENT_DATE,
    time_spent INTERVAL NOT NULL DEFAULT '0 hours 0 minutes',
    UNIQUE(habit_id, day)
);

CREATE TABLE IF NOT EXISTS habits_completed (
    id SERIAL PRIMARY KEY,
    habit_id INT NOT NULL REFERENCES habits(id) ON DELETE CASCADE,
    time_log_id INT NOT NULL REFERENCES habits_time_logs(id) ON DELETE CASCADE,
    completion_date TIMESTAMP NOT NULL DEFAULT NOW()
);
