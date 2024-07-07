-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS keys (
    key UUID NOT NULL,
    folder_id UUID NOT NULL REFERENCES folders (id) ON DELETE CASCADE,
    access_lvl INTEGER NOT NULL,
    PRIMARY KEY (key),
    UNIQUE (folder_id, access_lvl)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS keys;
-- +goose StatementEnd
