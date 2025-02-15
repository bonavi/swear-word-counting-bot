-- +goose Up
-- +goose StatementBegin
ALTER TABLE swears_counting_bot.messages DROP CONSTRAINT messages_pkey;
ALTER TABLE swears_counting_bot.messages RENAME COLUMN id TO tg_id;
ALTER TABLE swears_counting_bot.messages ADD id BIGINT GENERATED ALWAYS AS IDENTITY NOT NULL;
ALTER TABLE swears_counting_bot.messages ADD CONSTRAINT messages_pk PRIMARY KEY (id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE swears_counting_bot.messages DROP CONSTRAINT messages_pk;
ALTER TABLE swears_counting_bot.messages DROP COLUMN id;
ALTER TABLE swears_counting_bot.messages RENAME COLUMN tg_id TO id;
ALTER TABLE swears_counting_bot.messages ADD CONSTRAINT messages_pkey PRIMARY KEY (id);
-- +goose StatementEnd
