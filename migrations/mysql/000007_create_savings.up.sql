CREATE TABLE savings (
    id CHAR(36) DEFAULT (UUID()) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    saldo INT(11) NOT NULL,
    saldo_in INT(11),
    saldo_out INT(11),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT
);