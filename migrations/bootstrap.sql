CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    document_number BIGINT NOT NULL,
    phone BIGINT NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE Events (
    event_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    start_datetime TIMESTAMP NOT NULL,
    end_datetime TIMESTAMP NOT NULL,
    location TEXT,
    visibility TEXT, -- public, private, etc.
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);
CREATE TABLE Event_Participants (
    event_id INT NOT NULL,
    user_id INT NOT NULL,
    status TEXT, -- accepted, maybe, declined, etc.
    PRIMARY KEY (event_id, user_id),
    FOREIGN KEY (event_id) REFERENCES Events(event_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);
CREATE TABLE calendar_schedule (
    schedule_id SERIAL PRIMARY KEY,
    calendar_id INT,
    weekday INT,
    start_time INT,
    finish_time INT,
    FOREIGN KEY (calendar_id) REFERENCES Calendars(calendar_id) ON DELETE CASCADE
);

CREATE TABLE Calendars (
    calendar_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    calendar_name TEXT NOT NULL,
    color TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);
CREATE TABLE Event_Categories (
    category_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    category_name TEXT NOT NULL,
    color TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);

-- CREATE TABLE Reminders (
--     reminder_id SERIAL PRIMARY KEY,
--     event_id INT NOT NULL,
--     user_id INT NOT NULL,
--     remind_datetime TIMESTAMP NOT NULL,
--     reminder_message TEXT,
--     status TEXT, -- pending, sent, dismissed, etc.
--     FOREIGN KEY (event_id) REFERENCES Events(event_id) ON DELETE CASCADE,
--     FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
-- );