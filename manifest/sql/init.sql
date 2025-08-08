CREATE TABLE IF NOT EXISTS `users` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `cid` INT NOT NULL,  -- Customer/Company ID (đa tenant)
    `username` VARCHAR(50) NOT NULL,
    `email` VARCHAR(100) NOT NULL,
    `password_hash` VARCHAR(255) NOT NULL COMMENT 'hashed password',
    `role` VARCHAR(50) DEFAULT 'user' COMMENT 'user role: admin, agent, etc.',
    `status` TINYINT DEFAULT 1 COMMENT '0: inactive, 1: active, 2: suspended',
    `avatar` VARCHAR(255) DEFAULT NULL COMMENT 'profile picture URL',
    `last_login_at` TIMESTAMP NULL DEFAULT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT UC_Username UNIQUE (cid, username),
    CONSTRAINT UC_Email UNIQUE (cid, email)
);
