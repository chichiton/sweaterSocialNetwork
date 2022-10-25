create table sweater_db.user_friend
(
    user_id bigint not null,
    friend_id bigint not null,
    constraint friend_pk
        primary key (user_id, friend_id),
    constraint user_friend_user_user_id_fk
        foreign key (user_id) references user (id),
    constraint user_friend_user_friend_id_fk
        foreign key (friend_id) references user (id)
);