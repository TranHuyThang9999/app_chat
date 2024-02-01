-- Active: 1701521287160@@127.0.0.1@5432@app_chat

create Table users(
    id BIGINT PRIMARY KEY,
    user_name varchar(255),
    age int,
    address varchar(255),
    avatar varchar(512),
    email VARCHAR(255),
    created_at int
);
CREATE TABLE messages (
    id BIGINT PRIMARY KEY,
    content text,
    sender_id BIGINT,
    recipient_id BIGINT,
    created_at INT
);

CREATE Table rooms (
    id BIGINT PRIMARY KEY,
    sender_id BIGINT,
    recipient_id BIGINT,
    created_at INT
);

select *from users;

delete from users;

