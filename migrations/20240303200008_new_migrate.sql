-- +goose Up


create table chat (
    id                      BIGSERIAL PRIMARY KEY,
    name                    VARCHAR(50) NOT NULL UNIQUE,
    is_active               BOOL NOT NULL DEFAULT true
);

create table users (
    id BIGSERIAL PRIMARY KEY,
    login varchar(50) UNIQUE
);

create table chat_users_list (
    chat_id                 BIGSERIAL REFERENCES chat(id) NOT NULL,
    user_id                 BIGSERIAL REFERENCES users(id) NOT NULL
);


create table messages (
    id                      BIGSERIAL PRIMARY KEY,
    chat_id                 BIGSERIAL REFERENCES chat(id) NOT NULL,
    user_id                 BIGSERIAL REFERENCES users(id) NOT NULL,
    message                 VARCHAR(300) NOT NULL,
    create_at               TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);




-- +goose Down
drop table messages;
drop table chat_users_list;
drop table users;
drop table chat;

