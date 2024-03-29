CREATE user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid BIGINT NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(30),
    is_phone_actived BOOLEAN DEFAULT FALSE,
    avatar VARCHAR(255),
    phone_actived_at DATETIME,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
