create table user_balance (
    id serial primary key,
    amount float not null,
    transaction text not null,
);
