ALTER TABLE users
    ADD COLUMN tg_username varchar(500) unique,
    ADD COLUMN tg_chat_id bigint,
    ADD COLUMN gitlab_state_id int,
    ADD COLUMN gitlab_updated_at timestamp with time zone,
    ADD COLUMN gitlab_website_url varchar(500);

---- create above / drop below ----

ALTER TABLE users
    DROP COLUMN tg_username,
    DROP COLUMN tg_chat_id,
    DROP COLUMN gitlab_state_id,
    DROP COLUMN gitlab_updated_at,
    DROP COLUMN gitlab_website_url;
