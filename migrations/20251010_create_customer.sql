-- Table: customer_type
CREATE TABLE customer_type (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    type VARCHAR(255) NOT NULL,
    status int(2)NOT NULL default 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_by BIGINT,
    updated_at DATETIME NULL,
    updated_by BIGINT NULL,
    deleted_at DATETIME NULL,
    deleted_by BIGINT NULL
);

-- Table: customer
CREATE TABLE customer (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status int(2)NOT NULL default 1,
    customer_type_id BIGINT,
    pharmacist VARCHAR(255),
    sipa_number VARCHAR(255),
    sipa_expired_at DATETIME NULL,
    address TEXT,
    phone VARCHAR(50),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_by BIGINT,
    updated_at DATETIME NULL,
    updated_by BIGINT NULL,
    deleted_at DATETIME NULL,
    deleted_by BIGINT NULL,
    FOREIGN KEY (customer_type_id) REFERENCES customer_type(id)
);
