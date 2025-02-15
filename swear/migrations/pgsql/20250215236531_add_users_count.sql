
-- +goose Up
-- +goose StatementBegin
ALTER TABLE swears_counting_bot.chats ADD COLUMN count_users INT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE swears_counting_bot.chats DROP COLUMN count_users;
-- +goose StatementEnd
