-- +migrate Up
create table todo_items (
    id bigserial primary key,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    deleted_at timestamp with time zone,
    uuid uuid unique not null,
    description varchar(255) not null,
    due_date timestamp not null
);

create index idx_todo_items_due_date on todo_items (due_date);

-- +migrate Down
drop table if exists todo_items;
