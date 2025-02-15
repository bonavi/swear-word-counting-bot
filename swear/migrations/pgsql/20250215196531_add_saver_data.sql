
-- +goose Up
-- +goose StatementBegin

CREATE TABLE swears_counting_bot.chats (
    id BIGINT PRIMARY KEY,
    type VARCHAR(255),
    title VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    username VARCHAR(255),
    bio TEXT,
    description TEXT,
    invite_link TEXT,
    slow_mode INT,
    sticker_set VARCHAR(255),
    can_set_sticker_set BOOLEAN,
    custom_emoji_set_name VARCHAR(255),
    linked_chat_id BIGINT,
    private BOOLEAN,
    protected BOOLEAN,
    no_voice_and_video BOOLEAN,
    has_hidden_members BOOLEAN,
    aggressive_anti_spam BOOLEAN,
    custom_emoji_id VARCHAR(255),
    emoji_expiration_unixtime BIGINT,
    background_emoji_id VARCHAR(255),
    accent_color_id INT,
    profile_accent_color_id INT,
    profile_background_emoji_id VARCHAR(255),
    has_visible_history BOOLEAN,
    unrestrict_boosts INT
);

CREATE TABLE swears_counting_bot.messages (
    id INT PRIMARY KEY,
    thread_id INT,
    unixtime BIGINT,
    original_message_id INT,
    original_signature VARCHAR(255),
    original_sender_name VARCHAR(255),
    original_unixtime BIGINT,
    automatic_forward BOOLEAN,
    last_edit BIGINT,
    topic_message BOOLEAN,
    protected BOOLEAN,
    album_id VARCHAR(255),
    signature VARCHAR(255),
    text TEXT,
    payload TEXT,
    caption TEXT,
    new_group_title VARCHAR(255),
    group_photo_deleted BOOLEAN,
    group_created BOOLEAN,
    super_group_created BOOLEAN,
    channel_created BOOLEAN,
    migrate_to BIGINT,
    migrate_from BIGINT,
    connected_website VARCHAR(255),
    sender_boost_count INT,
    has_media_spoiler BOOLEAN
);

CREATE TABLE swears_counting_bot.users (
    id BIGINT PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    is_forum BOOLEAN,
    username VARCHAR(255),
    language_code VARCHAR(10),
    is_bot BOOLEAN,
    is_premium BOOLEAN,
    added_to_menu BOOLEAN,
    usernames TEXT[],
    custom_emoji_status VARCHAR(255)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE swears_counting_bot.chats;
DROP TABLE swears_counting_bot.messages;
DROP TABLE swears_counting_bot.users;
-- +goose StatementEnd
