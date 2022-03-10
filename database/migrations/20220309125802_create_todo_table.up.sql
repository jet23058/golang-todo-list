CREATE TABLE IF NOT EXISTS todos(
    id serial PRIMARY KEY,
    uuid VARCHAR(36) UNIQUE NOT NULL,
    title VARCHAR (255) NOT NULL,
    content VARCHAR (255) NOT NULL,
    status VARCHAR (10) NOT NULL,
    created_by VARCHAR (255) NULL,
    updated_by VARCHAR (255) NULL,
    deleted_by VARCHAR (255) NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NULL,
    deleted_at timestamp NULL
);