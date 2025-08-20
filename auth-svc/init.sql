-- Auth Service Database Initialization
-- This script contains tables related to authentication, authorization, and user management

-- Create database for auth service
CREATE DATABASE IF NOT EXISTS auth_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE auth_db;

-- Independent tables (no foreign key dependencies)
CREATE TABLE invalidated_tokens (
    id VARCHAR(255) PRIMARY KEY,
    expiry_time DATETIME NOT NULL,
    INDEX idx_token_expiry (expiry_time)
) ENGINE=InnoDB;

CREATE TABLE permission (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL UNIQUE,
    INDEX idx_permission_name (name)
) ENGINE=InnoDB;

CREATE TABLE role (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL UNIQUE,
    INDEX idx_role_name (name)
) ENGINE=InnoDB;

-- Tables with foreign key dependencies
CREATE TABLE comment (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    content TEXT NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('APPROVED', 'PENDING', 'REJECTED')),
    movie_id BIGINT,
    -- Note: movie_id references movie table in cinemahouse_db
    -- This will be handled via service communication
    INDEX idx_comment_movie (movie_id)
) ENGINE=InnoDB;

CREATE TABLE user (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    gender VARCHAR(50) NOT NULL DEFAULT 'MALE' CHECK (gender IN ('MALE', 'FEMALE', 'OTHER')),
    phone_number VARCHAR(15) NOT NULL,
    dob DATE,
    otp VARCHAR(6),
    otp_expiry_date DATETIME,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    account VARCHAR(50) NOT NULL DEFAULT 'ACTIVE' CHECK (account IN ('ACTIVE', 'INACTIVE', 'BANNED')),
    comment_id BIGINT,
    FOREIGN KEY (comment_id) REFERENCES comment(id) ON DELETE SET NULL,
    INDEX idx_user_email (email),
    INDEX idx_user_username (username)
) ENGINE=InnoDB;

-- Junction tables for many-to-many relationships
CREATE TABLE user_roles (
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE,
    INDEX idx_user_roles_user (user_id),
    INDEX idx_user_roles_role (role_id)
) ENGINE=InnoDB;

CREATE TABLE role_permissions (
    role_id BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,
    PRIMARY KEY (role_id, permission_id),
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permission(id) ON DELETE CASCADE,
    INDEX idx_role_permissions_role (role_id),
    INDEX idx_role_permissions_perm (permission_id)
) ENGINE=InnoDB;

-- Insert default permissions
INSERT INTO permission (name) VALUES 
('USER_READ'),
('USER_WRITE'),
('USER_DELETE'),
('ROLE_READ'),
('ROLE_WRITE'),
('ROLE_DELETE'),
('PERMISSION_READ'),
('PERMISSION_WRITE'),
('PERMISSION_DELETE'),
('COMMENT_READ'),
('COMMENT_WRITE'),
('COMMENT_DELETE'),
('COMMENT_APPROVE');

-- Insert default roles
INSERT INTO role (name) VALUES 
('ADMIN'),
('USER'),
('MODERATOR');

-- Assign permissions to ADMIN role (all permissions)
INSERT INTO role_permissions (role_id, permission_id) 
SELECT r.id, p.id 
FROM role r, permission p 
WHERE r.name = 'ADMIN';

-- Assign basic permissions to USER role
INSERT INTO role_permissions (role_id, permission_id) 
SELECT r.id, p.id 
FROM role r, permission p 
WHERE r.name = 'USER' 
AND p.name IN ('USER_READ', 'COMMENT_READ', 'COMMENT_WRITE');

-- Assign moderate permissions to MODERATOR role
INSERT INTO role_permissions (role_id, permission_id) 
SELECT r.id, p.id 
FROM role r, permission p 
WHERE r.name = 'MODERATOR' 
AND p.name IN ('USER_READ', 'COMMENT_READ', 'COMMENT_WRITE', 'COMMENT_DELETE', 'COMMENT_APPROVE');
