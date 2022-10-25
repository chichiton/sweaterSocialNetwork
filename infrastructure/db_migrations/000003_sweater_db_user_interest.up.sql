create table sweater_db.user_interest
(
    user_id bigint not null ,
    title varchar(100) not null,
    constraint interest_pk
        primary key (user_id, title),
    constraint user_interest_user_fk
        foreign key (user_id) references user (id)
);