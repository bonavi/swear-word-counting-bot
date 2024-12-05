-- +goose Up
-- +goose StatementBegin

-- Создаем схему
CREATE SCHEMA swears_counting_bot;

-- Создаем таблицу ругательств
CREATE TABLE swears_counting_bot.swears (
	"text" varchar NOT NULL,
	CONSTRAINT swears_pk PRIMARY KEY ("text")
);

-- Создаем таблицу статистики
CREATE TABLE swears_counting_bot."statistics" (
	id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
	message_id BIGINT NOT NULL,
	chat_id BIGINT NOT NULL,
	user_id BIGINT NOT NULL,
	swear varchar NOT NULL,
	CONSTRAINT statistics_pk PRIMARY KEY (id),
	CONSTRAINT statistics_swears_fk FOREIGN KEY (swear) REFERENCES swears_counting_bot.swears("text")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA swears_counting_bot
-- +goose StatementEnd
