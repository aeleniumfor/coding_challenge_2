CREATE DATABASE webapp;

\c webapp;

CREATE TABLE users (
  id         INTEGER,
  name       TEXT,
  email      TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);