-------------------------------------------------------------------------------------

CREATE TYPE GENDER AS ENUM('MALE', 'FEMALE', 'OTHER');
CREATE TYPE EXPERTIZE_LEVEL AS ENUM('EXPERT', 'PROFICIENT', 'COMPETENT', 'ADVANCED_BEGINNER', 'BEGINNER');
CREATE TYPE EXPERTIZE_AREA AS ENUM('DENTISTRY', 'NUTRITION', 'PATHOLOGY', 'RADIOLOGY', 'SURGERY', 'TOXICOLOGY');

-------------------------------------------------------------------------------------

CREATE TABLE Account (
    account_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    phone_number VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    gender GENDER NOT NULL,
    join_date TIMESTAMP NOT NULL
);

CREATE TABLE Address(
    address_id SERIAL PRIMARY KEY,
    account_id INT REFERENCES Account, -- FK
    country VARCHAR(25) NOT NULL,
    city VARCHAR(25) NOT NULL,
    street VARCHAR(25) NOT NULL,
    plaque VARCHAR(25) NOT NULL
);

CREATE TABLE Vet(
    expertize_level EXPERTIZE_LEVEL NOT NULL,
    expertize_area EXPERTIZE_AREA NOT NULL,
    working_hours_start DATE NOT NULL,
    working_hours_end DATE NOT NULL
) INHERITS (Account);

CREATE TABLE Pet_Lover(
    free_time_start DATE NOT NULL,
    free_time_end DATE NOT NULL
) INHERITS (Account);

CREATE TABLE Pet_Owner();

CREATE TABLE Shop_Owner();


CREATE TABLE Address(
    address_id SERIAL PRIMARY KEY,
    account_id INT REFERENCES Account, -- FK
    country VARCHAR(25) NOT NULL,
    city VARCHAR(25) NOT NULL,
    street VARCHAR(25) NOT NULL,
    plaque VARCHAR(25) NOT NULL
);

-------------------------------------------------------------------------------------

CREATE TABLE Post(
    post_id SERIAL PRIMARY KEY,
    account_id INT REFERENCES Account, -- FK
    description TEXT NOT NULL
);

CREATE TABLE Post_Comment(
    comment_id SERIAL PRIMARY KEY,
    post_id INT REFERENCES Post, -- FK
    comment_text TEXT NOT NULL
);

-------------------------------------------------------------------------------------

CREATE TABLE Notification (
    notification_id SERIAL PRIMARY KEY,
    account_id INT REFERENCES Account, -- FK
    description TEXT NOT NULL,
    notify_time TIMESTAMP NOT NULL
);

CREATE TABLE Calendar(
    calendar_id SERIAL PRIMARY KEY,
    account_id INT REFERENCES Account, -- FK
    task_title VARCHAR(25) NOT NULL,
    task_time_start DATE NOT NULL,
    task_time_end DATE NOT NULL
);

CREATE TABLE Reaction (
    comment_id INT,
    account_id INT,
    PRIMARY KEY (account_id, comment_id),
    FOREIGN KEY (account_id) REFERENCES Account,
    FOREIGN KEY (comment_id) REFERENCES Post_Comment,
    up_vote BOOLEAN DEFAULT FALSE,
    down_vote BOOLEAN DEFAULT FALSE
);

-------------------------------------------------------------------------------------


CREATE TABLE Shop(
    shop_id SERIAL PRIMARY KEY,
    shop_owner_id INT REFERENCES Shop_Owner,
    shop_name VARCHAR(25) NOT NULL
);

CREATE TABLE Pet(
    pet_id SERIAL PRIMARY KEY,
    shop_id INT REFERENCES Shop,
    name VARCHAR(25),
    age INT NOT NULL,
    disability BOOLEAN NOT NULL
);

CREATE TABLE Product(
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    price FLOAT NOT NULL,
    description TEXT
);

CREATE TABLE Shop_Product(
    shop_id INT,
    product_id INT,
    PRIMARY KEY (shop_id, product_id),
    FOREIGN KEY (shop_id) REFERENCES Shop,
    FOREIGN KEY (product_id) REFERENCES Product
);
