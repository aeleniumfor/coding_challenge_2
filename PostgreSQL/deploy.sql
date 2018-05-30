CREATE DATABASE webapp;

\c webapp;

CREATE TABLE users (
  id         SERIAL PRIMARY KEY,
  name       TEXT,
  email      TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);