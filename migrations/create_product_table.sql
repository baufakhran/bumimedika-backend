

CREATE TABLE products (
    id              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name            VARCHAR(255) NOT NULL,
    batch_number    VARCHAR(100) NOT NULL,
    purchase_price  DECIMAL(15,2) NOT NULL,
    selling_price   DECIMAL(15,2) NOT NULL,
    stock           INT UNSIGNED DEFAULT 0,
    sold            INT UNSIGNED DEFAULT 0,
    unit            VARCHAR(50)  NOT NULL DEFAULT '',
    category        VARCHAR(100) NOT NULL DEFAULT '',
    location        VARCHAR(100) NOT NULL DEFAULT '',
    color           VARCHAR(50)  NOT NULL DEFAULT '',
    size            VARCHAR(50)  NOT NULL DEFAULT '',
    brand           VARCHAR(100) NOT NULL DEFAULT '',
    expired_at      DATE,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

