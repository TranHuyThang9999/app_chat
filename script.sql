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

CREATE TABLE message_status (
    id BIGINT PRIMARY KEY,
    message_id BIGINT,
    sender_id BIGINT,
    recipient_id BIGINT,
    is_sent BOOLEAN,--để đánh dấu xem tin nhắn đã được gửi thành công hay chưa.
    is_received BOOLEAN,-- để đánh dấu xem tin nhắn đã được người nhận nhận được hay chưa.
    created_at INT,
);

CREATE Table rooms (
    id BIGINT PRIMARY KEY,
    sender_id BIGINT,
    recipient_id BIGINT,
    created_at INT
);

select *from users;

delete from users;

SELECT * FROM "users" WHERE user_name = 'thang' ORDER BY "users"."id" LIMIT 1