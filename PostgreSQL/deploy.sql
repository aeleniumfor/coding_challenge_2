CREATE DATABASE webapp;

\c webapp;

CREATE TABLE users (
  id         SERIAL PRIMARY KEY,
  name       TEXT,
  email      TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);