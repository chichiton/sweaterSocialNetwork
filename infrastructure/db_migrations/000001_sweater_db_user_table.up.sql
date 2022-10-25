create table sweater_db.user
(
    id bigint auto_increment,
    login varchar(50) unique not null,
    password_hash varchar(100) not null,
    create_date datetime not null default now(),
    constraint user_pk
        primary key (id)
);
