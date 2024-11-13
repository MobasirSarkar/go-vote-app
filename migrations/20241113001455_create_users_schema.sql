-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE role_type AS ENUM ('user','admin');
CREATE TABLE users (
   user_id CHAR(8) PRIMARY KEY DEFAULT LEFT(REPLACE(uuid_generate_v4()::text,'-',''),8),
   name VARCHAR(100) NOT NULL,
   email VARCHAR(100) UNIQUE NOT NULL,
   role role_type NOT NULL DEFAULT 'user',
   password VARCHAR(100) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   delete_at TIMESTAMP DEFAULT NULL

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS "uuid-ossp";
DROP TYPE IF EXISTS role_type;
-- +goose StatementEnd
