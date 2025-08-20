-- Cinema House Service Database Initialization
-- This script contains tables related to cinema operations, movies, tickets, and bookings

-- Create database for cinemahouse service
CREATE DATABASE IF NOT EXISTS cinemahouse_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE cinemahouse_db;

-- Independent tables (no foreign key dependencies)
CREATE TABLE room_type (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL UNIQUE,
    status VARCHAR(50) NOT NULL CHECK (status IN ('ACTIVE', 'INACTIVE')),
    INDEX idx_room_type_name (name)
) ENGINE=InnoDB;

CREATE TABLE seat_type (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL UNIQUE,
    INDEX idx_seat_type_name (name)
) ENGINE=InnoDB;

CREATE TABLE movie_type (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    type VARCHAR(100) NOT NULL UNIQUE,
    movie_type_status VARCHAR(50) NOT NULL CHECK (movie_type_status IN ('ACTIVE', 'INACTIVE')),
    INDEX idx_movie_type (type)
) ENGINE=InnoDB;

CREATE TABLE ticket_price (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    price DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) CHECK (status IN ('ACTIVE', 'INACTIVE')),
    INDEX idx_ticket_price_price (price)
) ENGINE=InnoDB;

CREATE TABLE cinema (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    cinema_name VARCHAR(255) NOT NULL UNIQUE,
    cinema_address VARCHAR(255) NOT NULL,
    cinema_status VARCHAR(50) NOT NULL CHECK (cinema_status IN ('ACTIVE', 'INACTIVE')),
    INDEX idx_cinema_name (cinema_name)
) ENGINE=InnoDB;

CREATE TABLE movie (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    movie_name VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    duration TIME,
    language VARCHAR(100) NOT NULL,
    director VARCHAR(255) NOT NULL,
    actors VARCHAR(255) NOT NULL,
    age_limit TINYINT UNSIGNED NOT NULL,
    cover_photo TEXT NOT NULL,
    release_date DATE NOT NULL,
    start_date DATE NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('UPCOMING', 'RUNNING', 'ENDED')),
    INDEX idx_movie_name (movie_name),
    INDEX idx_movie_release (release_date)
) ENGINE=InnoDB;

CREATE TABLE payment (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    payment_method VARCHAR(50) NOT NULL CHECK (payment_method IN ('CREDIT_CARD', 'DEBIT_CARD', 'PAYPAL', 'BANK_TRANSFER', 'CASH')),
    payment_status VARCHAR(50) NOT NULL CHECK (payment_status IN ('PENDING', 'COMPLETED', 'FAILED', 'CANCELLED')),
    payment_date DATETIME NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    transaction_id VARCHAR(255) UNIQUE,
    INDEX idx_payment_status (payment_status),
    INDEX idx_payment_date (payment_date)
) ENGINE=InnoDB;

-- Tables with foreign key dependencies
CREATE TABLE screening_room (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    screening_room_status VARCHAR(50) NOT NULL CHECK (screening_room_status IN ('ACTIVE', 'INACTIVE')),
    roomtype_id BIGINT NOT NULL,
    cinema_id BIGINT NOT NULL,
    FOREIGN KEY (roomtype_id) REFERENCES room_type(id) ON DELETE RESTRICT,
    FOREIGN KEY (cinema_id) REFERENCES cinema(id) ON DELETE RESTRICT,
    INDEX idx_screening_room_name (name),
    INDEX idx_screening_room_type (roomtype_id),
    INDEX idx_screening_room_cinema (cinema_id)
) ENGINE=InnoDB;

CREATE TABLE screening_session (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    start_date DATE NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    screening_session_status VARCHAR(50) NOT NULL CHECK (screening_session_status IN ('SCHEDULED', 'ONGOING', 'ENDED')),
    movie_id BIGINT NOT NULL,
    screening_room_id BIGINT NOT NULL,
    FOREIGN KEY (movie_id) REFERENCES movie(id) ON DELETE RESTRICT,
    FOREIGN KEY (screening_room_id) REFERENCES screening_room(id) ON DELETE RESTRICT,
    INDEX idx_session_movie (movie_id),
    INDEX idx_session_room (screening_room_id)
) ENGINE=InnoDB;

CREATE TABLE seat (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    seat_row VARCHAR(10) NOT NULL,
    seat_column TINYINT UNSIGNED NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('AVAILABLE', 'RESERVED', 'SOLD')),
    screening_room_id BIGINT NOT NULL,
    seat_type_id BIGINT NOT NULL,
    FOREIGN KEY (screening_room_id) REFERENCES screening_room(id) ON DELETE CASCADE,
    FOREIGN KEY (seat_type_id) REFERENCES seat_type(id) ON DELETE RESTRICT,
    INDEX idx_seat_room (screening_room_id),
    INDEX idx_seat_type (seat_type_id)
) ENGINE=InnoDB;

CREATE TABLE bill (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    bill_date DATE NOT NULL,
    bill_time DATETIME NOT NULL,
    total_amount DECIMAL(15, 2) NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('PENDING', 'PAID', 'CANCELLED')),
    user_id BIGINT NOT NULL,
    -- Note: user_id references user table in auth_db
    -- This will be handled via service communication
    payment_id BIGINT,
    request_id VARCHAR(255),
    FOREIGN KEY (payment_id) REFERENCES payment(id) ON DELETE SET NULL,
    INDEX idx_bill_user (user_id),
    INDEX idx_bill_payment (payment_id)
) ENGINE=InnoDB;

CREATE TABLE ticket (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    status VARCHAR(50) NOT NULL CHECK (status IN ('AVAILABLE', 'SOLD', 'CANCELLED')),
    screening_session_id BIGINT NOT NULL,
    ticket_price_id BIGINT NOT NULL,
    bill_id BIGINT,
    seat_id BIGINT NOT NULL,
    FOREIGN KEY (screening_session_id) REFERENCES screening_session(id) ON DELETE RESTRICT,
    FOREIGN KEY (ticket_price_id) REFERENCES ticket_price(id) ON DELETE RESTRICT,
    FOREIGN KEY (bill_id) REFERENCES bill(id) ON DELETE SET NULL,
    FOREIGN KEY (seat_id) REFERENCES seat(id) ON DELETE RESTRICT,
    INDEX idx_ticket_session (screening_session_id),
    INDEX idx_ticket_price (ticket_price_id),
    INDEX idx_ticket_seat (seat_id)
) ENGINE=InnoDB;

-- Junction table for many-to-many relationship between movies and movie types
CREATE TABLE movie_typeofmovies (
    movie_id BIGINT NOT NULL,
    movie_type_id BIGINT NOT NULL,
    PRIMARY KEY (movie_id, movie_type_id),
    FOREIGN KEY (movie_id) REFERENCES movie(id) ON DELETE CASCADE,
    FOREIGN KEY (movie_type_id) REFERENCES movie_type(id) ON DELETE CASCADE,
    INDEX idx_movie_type_movie (movie_id),
    INDEX idx_movie_type_type (movie_type_id)
) ENGINE=InnoDB;

-- Insert default data
-- Insert default room types
INSERT INTO room_type (name, status) VALUES 
('Standard', 'ACTIVE'),
('Premium', 'ACTIVE'),
('VIP', 'ACTIVE'),
('IMAX', 'ACTIVE'),
('4DX', 'ACTIVE');

-- Insert default seat types
INSERT INTO seat_type (name) VALUES 
('Regular'),
('Premium'),
('VIP'),
('Recliner'),
('Couple');

-- Insert default movie types
INSERT INTO movie_type (type, movie_type_status) VALUES 
('Action', 'ACTIVE'),
('Comedy', 'ACTIVE'),
('Drama', 'ACTIVE'),
('Horror', 'ACTIVE'),
('Romance', 'ACTIVE'),
('Sci-Fi', 'ACTIVE'),
('Thriller', 'ACTIVE'),
('Animation', 'ACTIVE'),
('Documentary', 'ACTIVE'),
('Fantasy', 'ACTIVE');

-- Insert default ticket prices
INSERT INTO ticket_price (price, status) VALUES 
(50000.00, 'ACTIVE'),  -- Regular ticket
(75000.00, 'ACTIVE'),  -- Premium ticket
(100000.00, 'ACTIVE'), -- VIP ticket
(120000.00, 'ACTIVE'), -- IMAX ticket
(150000.00, 'ACTIVE'); -- 4DX ticket
