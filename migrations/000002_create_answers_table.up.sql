CREATE TABLE IF NOT EXISTS answers (
    id INT(10) UNSIGNED NOT NULL serial AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL default now(),
    updated_at DATETIME NOT NULL default now(),
    task_id INT(10) NOT NULL,
    is_correct BOOL NOT NULL DEFAULT FALSE,
    FOREIGN KEY (task_id) REFERENCES tasks(id),
    PRIMARY KEY (id)
    )ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE answers
    ALTER id ADD GENERATED ALWAYS AS IDENTITY;