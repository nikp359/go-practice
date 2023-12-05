CREATE TABLE user_visit (
      id SERIAL PRIMARY KEY,
      user_id INT NOT NULL,
      entered_at TIMESTAMP,
      exit_at TIMESTAMP
);