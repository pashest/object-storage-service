-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "pgcrypto"

CREATE TABLE IF NOT EXISTS storage_servers (
    server_uid  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    address     TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS storage_servers;
-- +goose StatementEnd
