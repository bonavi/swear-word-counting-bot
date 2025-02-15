
-- +goose Up
-- +goose StatementBegin

CREATE TABLE swears_counting_bot.chats (
    id BIGINT PRIMARY KEY NOT NULL,
    type VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    bio TEXT NOT NULL,
    description TEXT NOT NULL,
    invite_link TEXT NOT NULL,
    slow_mode INT NOT NULL,
    sticker_set VARCHAR(255) NOT NULL,
    can_set_sticker_set BOOLEAN NOT NULL,
    custom_emoji_set_name VARCHAR(255) NOT NULL,
    linked_chat_id BIGINT NOT NULL,
    private BOOLEAN NOT NULL,
    protected BOOLEAN NOT NULL,
    no_voice_and_video BOOLEAN NOT NULL,
    has_hidden_members BOOLEAN NOT NULL,
    aggressive_anti_spam BOOLEAN NOT NULL,
    custom_emoji_id VARCHAR(255) NOT NULL,
    emoji_expiration_unixtime BIGINT NOT NULL,
    background_emoji_id VARCHAR(255) NOT NULL,
    accent_color_id INT NOT NULL,
    profile_accent_color_id INT NOT NULL,
    profile_background_emoji_id VARCHAR(255) NOT NULL,
    has_visible_history BOOLEAN NOT NULL,
    unrestrict_boosts INT NOT NULL
);

CREATE TABLE swears_counting_bot.messages (
    id INT PRIMARY KEY NOT NULL,
    chat_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    thread_id INT NOT NULL,
    date_time timestamptz NOT NULL,
    original_message_id INT NOT NULL,
    original_signature VARCHAR(255) NOT NULL,
    original_sender_name VARCHAR(255) NOT NULL,
    original_unixtime BIGINT NOT NULL,
    automatic_forward BOOLEAN NOT NULL,
    last_edit BIGINT NOT NULL,
    topic_message BOOLEAN NOT NULL,
    protected BOOLEAN NOT NULL,
    album_id VARCHAR(255) NOT NULL,
    signature VARCHAR(255) NOT NULL,
    text TEXT NOT NULL,
    payload TEXT NOT NULL,
    caption TEXT NOT NULL,
    new_group_title VARCHAR(255) NOT NULL,
    group_photo_deleted BOOLEAN NOT NULL,
    group_created BOOLEAN NOT NULL,
    super_group_created BOOLEAN NOT NULL,
    channel_created BOOLEAN NOT NULL,
    migrate_to BIGINT NOT NULL,
    migrate_from BIGINT NOT NULL,
    connected_website VARCHAR(255) NOT NULL,
    sender_boost_count INT NOT NULL,
    has_media_spoiler BOOLEAN NOT NULL
);

CREATE TABLE swears_counting_bot.users (
    id BIGINT PRIMARY KEY NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    is_forum BOOLEAN NOT NULL,
    username VARCHAR(255) NOT NULL,
    language_code VARCHAR(10) NOT NULL,
    is_bot BOOLEAN NOT NULL,
    is_premium BOOLEAN NOT NULL,
    added_to_menu BOOLEAN NOT NULL,
    usernames TEXT[],
    custom_emoji_status VARCHAR(255) NOT NULL
);

ALTER TABLE swears_counting_bot.messages ADD CONSTRAINT messages_chats_fk FOREIGN KEY (chat_id) REFERENCES swears_counting_bot.chats(id);
ALTER TABLE swears_counting_bot.messages ADD CONSTRAINT messages_users_fk FOREIGN KEY (user_id) REFERENCES swears_counting_bot.users(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE swears_counting_bot.messages DROP CONSTRAINT messages_chats_fk;
ALTER TABLE swears_counting_bot.messages DROP CONSTRAINT messages_users_fk;

DROP TABLE swears_counting_bot.chats;
DROP TABLE swears_counting_bot.messages;
DROP TABLE swears_counting_bot.users;
-- +goose StatementEnd
