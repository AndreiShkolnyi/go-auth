-- +goose Up
create table auth (
    id serial primary key,
    name text not null,
    email text unique not null,
    password text not null,
    password_confirm text not null,
    role text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table auth;
