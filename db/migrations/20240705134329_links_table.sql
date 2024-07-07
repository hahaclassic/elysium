-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Links (
    url TEXT NOT NULL,
    tag TEXT NOT NULL,
    folder_id UUID NOT NULL REFERENCES folders (id) ON DELETE CASCADE,
    PRIMARY KEY (folder_id, url)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Links;
-- +goose StatementEnd
