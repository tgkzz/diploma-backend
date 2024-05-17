CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    fname VARCHAR(50),
    lname VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS admins (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE,
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS experts (
    id SERIAL PRIMARY KEY,
    fname VARCHAR(50),
    lname VARCHAR(50),
    email VARCHAR(255) UNIQUE,
    cost NUMERIC(10, 2),
    password VARCHAR(255),
    description TEXT
);

-- must be moved to mongo
-- CREATE TABLE IF NOT EXISTS course (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(255),
--     description TEXT
-- );

CREATE TABLE IF NOT EXISTS course_transactions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    course_id INTEGER,
    cost NUMERIC
);

CREATE TABLE IF NOT EXISTS meeting_transactions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    expert_id INTEGER,
    time_start TIMESTAMP,
    time_end TIMESTAMP,
    total_cost NUMERIC,
    meeting_link TEXT
);

