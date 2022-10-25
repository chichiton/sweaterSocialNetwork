create table sweater_db.user_profile
(
    user_id bigint not null,
    first_name varchar(50) not null,
    last_name  varchar(50) not null,
    age int not null,
    gender int not null,
    city varchar(50) not null,
    constraint user_profile_pk
        primary key (user_id),
    constraint user_profile_user_fk
        foreign key (user_id) references user (id)
);