create table user_balance (
    id uuid not null,
    time timestamp with time zone not null,
    state text not null,
    operation_amount float not null,
    balance float not null,
    transaction text not null
);

create index transaltion_idx on user_balance (transaction);

insert into user_balance values (null, current_timestamp, null, null, 0, null)
