CREATE TABLE tasks
(
    id serial primary key,
    title varchar(255) not null,
    description varchar(255),
    created_at timestamp not null,
    updated_at timestamp
);