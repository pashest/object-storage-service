-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS storage_servers (
    address  TEXT NOT NULL PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS storage_servers;
-- +goose StatementEnd
