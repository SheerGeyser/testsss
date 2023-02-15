ALTER TABLE users
    ADD COLUMN gitlab_username varchar(500);

CREATE UNIQUE INDEX users_gitlab_username_uindex
    ON users (gitlab_username);

---- create above / drop below ----

ALTER TABLE users
    DROP COLUMN gitlab_username;
