CREATE TABLE IF NOT EXISTS todos(  
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    activity_group_id INT,
    title VARCHAR(100),
    is_active TINYINT(1) DEFAULT 1,
    priority VARCHAR(100) DEFAULT 'very-high',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL
) DEFAULT CHARSET UTF8;