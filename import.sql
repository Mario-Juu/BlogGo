CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT,
    email varchar(255) UNIQUE,
    password varchar(255),
    PRIMARY KEY (id)
);

INSERT INTO users(email, password) values ("mariojralves2006@gmail.com", "1234");

CREATE TABLE posts(
    id int NOT NULL AUTO_INCREMENT,
    title varchar(255) NOT NULL,
    slug varchar (255) NOT NULL UNIQUE,
    content text,
    user_id int NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP(),
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP(),
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
