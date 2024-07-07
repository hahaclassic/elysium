-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_folders (
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    folder_id UUID REFERENCES folders (id) ON DELETE CASCADE,
    access_lvl INTEGER,
    PRIMARY KEY (user_id, folder_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_folders;
-- +goose StatementEnd
