CREATE TABLE user (
    user_id             INT AUTO_INCREMENT  PRIMARY KEY,
    email_address       VARCHAR(100) NOT NULL   UNIQUE,
    password_hash       VARCHAR(255) NOT NULL,
    created_timestamp   TIMESTAMP   NOT NULL    DEFAULT (current_timestamp()),
    modified_timestamp  TIMESTAMP   NOT NULL    DEFAULT (current_timestamp())
);

CREATE TRIGGER user_modified_timestamp_trigger
BEFORE UPDATE ON user
FOR EACH ROW
   SET NEW.modified_timestamp = current_timestamp();

CREATE TABLE category (
    category_id         INT AUTO_INCREMENT  PRIMARY KEY,
    user_id             INT NOT NULL,
    category_name       VARCHAR(255)    NOT NULL,
    created_timestamp   TIMESTAMP   NOT NULL    DEFAULT (current_timestamp()),
    modified_timestamp  TIMESTAMP   NOT NULL    DEFAULT (current_timestamp()),
    CONSTRAINT `fk_category_user_id`
        FOREIGN KEY (user_id) REFERENCES user (user_id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION
);

CREATE TRIGGER category_modified_timestamp_trigger
BEFORE UPDATE ON category
FOR EACH ROW
   SET NEW.modified_timestamp = current_timestamp();

CREATE TABLE task (
    task_id             INT AUTO_INCREMENT  PRIMARY KEY,
    user_id             INT NOT NULL,
    category_id         INT NOT NULL,
    task_name           VARCHAR(255) NOT NULL,
    task_description    TEXT,
    task_status         INT NOT NULL    DEFAULT 0   COMMENT '0 - Pending, 1 - In Progress, 2 - Completed',
    due_date            DATE    NOT NULL    DEFAULT (current_date()),
    created_timestamp   TIMESTAMP   NOT NULL    DEFAULT (current_timestamp()),
    modified_timestamp  TIMESTAMP   NOT NULL    DEFAULT (current_timestamp()),
    CONSTRAINT `fk_task_user_id`
        FOREIGN KEY (user_id) REFERENCES user (user_id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT `fk_task_category_id`
        FOREIGN KEY (category_id) REFERENCES category (category_id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION
);

CREATE TRIGGER task_modified_timestamp_trigger
BEFORE UPDATE ON task
FOR EACH ROW
   SET NEW.modified_timestamp = current_timestamp();
