create table user_balance (
    id uuid not null,
    time timestamp with time zone not null,
    state text not null,
    amount float not null,
    balance float not null,
    transaction text not null
);

create index transaltion_idx on user_balance (transaction);

insert into user_balance values ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', current_timestamp, '' , 0 , 0 , '');
