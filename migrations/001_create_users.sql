-- This is a sample migration.

create table users(
  id serial primary key,
  gitlab_id int,
  is_subscribed boolean not null default false
);

---- create above / drop below ----

drop table users;
