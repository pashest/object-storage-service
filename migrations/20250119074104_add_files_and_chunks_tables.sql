-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "pgcrypto"

CREATE TABLE IF NOT EXISTS file_info (
    file_uid   UUID DEFAULT gen_random_uuid()
    file_name  TEXT NOT NULL,
    user_id    TEXT NOT NULL,
    file_size  BIGINT NOT NULL,
    status     TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ  DEFAULT now() NOT NULL,

    PRIMARY KEY (file_name, user_id)
);

CREATE TABLE IF NOT EXISTS chunk_info (
    chunk_name     TEXT NOT NULL PRIMARY KEY,
    file_uid       UUID NOT NULL,
    chunk_number   SMALLINT NOT NULL,
    chunk_size     BIGINT NOT NULL,
    server_address TEXT NOT NULL,
    created_at     TIMESTAMPTZ DEFAULT now() NOT NULL,

    FOREIGN KEY (file_uid) REFERENCES file_info (file_uid),
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS files;
DROP TABLE IF EXISTS chunks;
-- +goose StatementEnd
