-- +migration Up

create table todo_items (
    id bigserial primary key,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    uuid uuid unique ,
    description varchar(255) not null,
    due_data timestamp not null
);

create index idx_todo_items_due_date on todo_items (due_data);

-- +migration Down

drop table if exists todo_items;
