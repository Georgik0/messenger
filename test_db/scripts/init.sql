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

INSERT INTO my_user (username)
VALUES
       ('test1'),
       ('test2'),
       ('test3');

INSERT INTO chat(name)
VALUES
    ('chat1'),
    ('chat2');

INSERT INTO chat_user(user_id, chat_id)
VALUES
    (1,1),
    (2,1),
    (1,2),
    (3,2);

INSERT INTO message(chat_id, author_id, text)
VALUES
    (1,1,'test1 write in chat1'),
    (2,1, 'test1 write in chat2');