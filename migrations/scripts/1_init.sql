-- +migrate Up
CREATE TABLE sessions
(
    id          BIGSERIAL,
    uuid        uuid                        NOT NULL,
    user_uuid   uuid PRIMARY KEY,
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    expiring_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TYPE gender_type AS ENUM ('male','female');

CREATE TABLE users
(
    id            BIGSERIAL,
    uuid          uuid PRIMARY KEY,
    version       INTEGER                     NOT NULL DEFAULT 1,
    created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    updated_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    deleted_at    TIMESTAMP WITHOUT TIME ZONE,
    first_name    varchar(256),
    last_name     varchar(256),
    bio           text,
    age           integer,
    gender        gender_type,
    interests     text,
    birth_date    date,
    email         varchar(512)                NOT NULL UNIQUE,
    password_hash text
);

CREATE TABLE friendships
(
    id             BIGSERIAL,
    uid            uuid PRIMARY KEY,
    version        INTEGER                     NOT NULL DEFAULT 1,
    created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    updated_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    deleted_at     TIMESTAMP WITHOUT TIME ZONE,
    requester_uuid uuid                        NOT NULL,
    receiver_uuid  uuid                        NOT NULL,
    accepted       BOOLEAN                     NOT NULL DEFAULT FALSE
);

CREATE UNIQUE INDEX friendships_unique ON friendships (requester_uuid, receiver_uuid);

-- +migrate Down
DROP TABLE sessions;