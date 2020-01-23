create table user_balance (
    id uuid not null,
    time timestamp with time zone not null,
    state text not null,
    value float not null,
    amount float not null,
    transaction text not null
);
