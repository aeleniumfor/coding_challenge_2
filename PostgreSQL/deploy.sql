CREATE DATABASE webapp;

\c webapp;

CREATE TABLE users (
  id         SERIAL PRIMARY KEY,
  name       TEXT,
  email      TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);

-- CREATE FUNCTION set_update_time()
--   RETURNS OPAQUE AS '
-- BEGIN
--   RETURN now();
-- END;
-- '
-- LANGUAGE 'plpgsql';
--
-- CREATE TRIGGER update_tri
--   BEFORE UPDATE -- only update
--   ON users
--   FOR EACH ROW
-- EXECUTE PROCEDURE set_update_time();