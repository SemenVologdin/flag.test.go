CREATE TABLE IF NOT EXISTS answers (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    task_id INT(10) NOT NULL,
    FOREIGN KEY (task_id) REFERENCES tasks(id),
    PRIMARY KEY (id)
    )ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;