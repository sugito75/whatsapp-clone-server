CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    phone CHAR(13) NOT NULL UNIQUE,
    password VARCHAR(250) NOT NULL,
    profile_picture TEXT,
    bio VARCHAR(250),
    is_online BOOLEAN DEFAULT FALSE,
    last_online TIMESTAMP DEFAULT NOW()
);

CREATE TABLE sessions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    session_id VARCHAR(250) NOT NULL,

    UNIQUE(user_id, session_id)
);

CREATE TABLE chats (
    id BIGSERIAL PRIMARY KEY,
    chat_type VARCHAR(10) CHECK (chat_type IN ('private', 'group')) NOT NULL,
    title VARCHAR(100),
    description TEXT,
    icon TEXT,
    last_read_message_id BIGINT DEFAULT NULL,
    created_by BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE chat_members (
    id BIGSERIAL PRIMARY KEY,
    chat_id BIGINT NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(10) DEFAULT 'member' CHECK(role IN ('member', 'admin')),
    joined_at TIMESTAMP DEFAULT NOW(),

    UNIQUE(chat_id, user_id)

);

CREATE TABLE messages (
    id BIGSERIAL PRIMARY KEY,
    chat_id BIGINT REFERENCES chats(id) ON DELETE CASCADE,
    sender_id INT REFERENCES users(id) ON DELETE SET NULL,
    content TEXT,
    reply_to BIGINT REFERENCES messages(id) ON DELETE SET NULL,
    sent_at TIMESTAMP DEFAULT NOW(),
    edited_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE message_statuses (
    id BIGSERIAL PRIMARY KEY,
    message_id BIGINT REFERENCES messages(id) ON DELETE CASCADE,
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    status VARCHAR(10) DEFAULT 'sent' CHECK(status IN ('sent', 'delivered', 'readed')),
    updated_at TIMESTAMP DEFAULT NOW(),

    UNIQUE(message_id, user_id)
);

CREATE TABLE statuses (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    media TEXT NOT NULL,
    media_type VARCHAR(10) DEFAULT 'text' CHECK (media_type IN ('text', 'image', 'video')),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE call_histories (
    id BIGSERIAL PRIMARY KEY,
    caller BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    called BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    duration INT DEFAULT 0, -- in seconds
    created_at TIMESTAMP DEFAULT NOW(),

    UNIQUE(caller, called)
);