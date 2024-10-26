-- +goose Up
CREATE TABLE IF NOT EXISTS "user"
(
    id    uuid         not null primary key,
    login varchar(255) not null unique,
    pass  varchar(255) not null
);

CREATE TYPE secret_kind AS ENUM
(
    'credentials',
    'credit_card',
    'text',
    'binary'
);

CREATE TABLE IF NOT EXISTS "secret"
(
    id      uuid         not null primary key,
    user_id uuid         not null REFERENCES "user" (id) ON DELETE CASCADE,
    name    varchar(255) not null,
    kind    secret_kind  not null,
    unique (user_id, name)
);

CREATE TABLE IF NOT EXISTS "tag"
(
    id        uuid         not null primary key,
    secret_id uuid         not null REFERENCES secret (id) ON DELETE CASCADE,
    name      varchar(255) not null,
    value     text         not null
);

CREATE TABLE IF NOT EXISTS "credentials"
(
    id       uuid         not null primary key REFERENCES secret (id) ON DELETE CASCADE,
    login    varchar(255) not null,
    password varchar(255) not null
);

CREATE TABLE IF NOT EXISTS "credit_card"
(
    id   uuid         not null primary key REFERENCES secret (id) ON DELETE CASCADE,
    date char(5)      not null,
    name varchar(255) not null,
    code char(3)      not null
);

CREATE TABLE IF NOT EXISTS "binary_data"
(
    id          uuid not null primary key REFERENCES secret (id) ON DELETE CASCADE,
    is_uploaded bool not null default false
);

-- +goose Down
DROP TABLE "binary_data";
DROP TABLE "credit_card";
DROP TABLE "credentials";
DROP TABLE "tag";
DROP TABLE "secret";
DROP TABLE "user";
DROP TYPE "secret_kind";
