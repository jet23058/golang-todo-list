CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   uuid VARCHAR(36) UNIQUE NOT NULL,
   name VARCHAR (255) UNIQUE NOT NULL,
   password VARCHAR (255) NOT NULL,
   email VARCHAR (255) UNIQUE NOT NULL,
   status VARCHAR (10) NOT NULL,
   created_by VARCHAR (255) NULL,
   updated_by VARCHAR (255) NULL,
   deleted_by VARCHAR (255) NULL,
   created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
   updated_at timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NULL,
   deleted_at timestamp NULL
);