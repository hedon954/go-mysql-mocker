CREATE TABLE IF NOT EXISTS `users` (
   `id` INT NOT NULL AUTO_INCREMENT,
   `username` VARCHAR(50) NOT NULL,
    `email` VARCHAR(100) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `users` (`username`, `email`) VALUES
('john_doe', 'john.doe@example.com'),
('jane_smith', 'jane.smith@example.com'),
('alice_jones', 'alice.jones@example.com');