CREATE TABLE my_user (
    id serial PRIMARY KEY,
    username varchar UNIQUE NOT NULL,
    created_at timestamp default current_timestamp
);

CREATE TABLE chat (
    id serial PRIMARY KEY,
    name varchar UNIQUE NOT NULL,
    created_at timestamp default current_timestamp
);

CREATE TABLE chat_user (
    user_id integer not null,
    chat_id integer not null,
    foreign key (user_id) references my_user(id),
    foreign key (chat_id) references chat(id)
);

CREATE TABLE message (
     id serial PRIMARY KEY,
     chat_id integer not null,
     foreign key (chat_id) references chat(id),
     author_id integer not null,
     foreign key (author_id) references my_user(id),
     text varchar,
     created_at timestamp default current_timestamp
);

/*CREATE TABLE chat_message
(
    id         integer PRIMARY KEY,
    chat_id integer not null,
    foreign key (chat_id) references chat(id),
    message_id integer not null,
    foreign key (message_id) references message(id)
);*/

/*CREATE TABLE user_message (
    id integer PRIMARY KEY,
    message_id integer not null,
    foreign key (message_id) references message(id)
);*/