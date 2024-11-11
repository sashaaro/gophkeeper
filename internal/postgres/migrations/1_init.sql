-- +goose Up
CREATE TABLE IF NOT EXISTS "user"
(
    id    uuid         not null primary key,
    login varchar(255) not null unique,
    pass  varchar(255) not null
);

CREATE TABLE IF NOT EXISTS "secret"
(
    id      uuid         not null primary key,
    user_id uuid         not null REFERENCES "user" (id) ON DELETE CASCADE,
    name    varchar(255) not null,
    value    bytea  not null,
    unique (user_id, name)
);


-- +goose Down
DROP TABLE "secret";
DROP TABLE "user";
