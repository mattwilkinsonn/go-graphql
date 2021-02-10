CREATE TABLE users(
  id serial primary key,
  username varchar not null unique,
  email varchar not null unique,
  password varchar not null,
);
---- create above / drop below ----
drop table users;