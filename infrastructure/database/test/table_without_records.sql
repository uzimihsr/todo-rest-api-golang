DROP DATABASE IF EXISTS todo_db;
CREATE DATABASE IF NOT EXISTS todo_db;
USE todo_db;

DROP TABLE IF EXISTS todo;
CREATE TABLE IF NOT EXISTS todo (
  id INT AUTO_INCREMENT PRIMARY KEY, 
  title VARCHAR(100),
  done BOOLEAN DEFAULT false,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);