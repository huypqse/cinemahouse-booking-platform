CREATE TABLE bill (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    bill_date DATE NOT NULL,
    bill_time DATETIME NOT NULL,
    total_amount DECIMAL(15, 2) NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('PENDING', 'PAID', 'CANCELLED')),
    user_id BIGINT NOT NULL,
    payment_id BIGINT UNIQUE,
    request_id VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT,
    FOREIGN KEY (payment_id) REFERENCES payment(id) ON DELETE SET NULL,
    INDEX idx_bill_user (user_id)
);

CREATE TABLE cinema (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    cinema_name VARCHAR(255) NOT NULL UNIQUE,
    cinema_address VARCHAR(255) NOT NULL,
    cinema_status VARCHAR(50) NOT NULL CHECK (cinema_status IN ('ACTIVE', 'INACTIVE')),
    INDEX idx_cinema_name (cinema_name)
);

CREATE TABLE comment (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    content TEXT NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('APPROVED', 'PENDING', 'REJECTED')),
    movie_id BIGINT,
    FOREIGN KEY (movie_id) REFERENCES movie(id) ON DELETE SET NULL,
    INDEX idx_comment_movie (movie_id)
);

CREATE TABLE invalidated_tokens (
    id VARCHAR(255) PRIMARY KEY,
    expiry_time DATETIME NOT NULL,
    INDEX idx_token_expiry (expiry_time)
);

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
);

CREATE TABLE movie_type (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    type VARCHAR(100) NOT NULL,
    movie_type_status VARCHAR(50) NOT NULL CHECK (movie_type_status IN ('ACTIVE', 'INACTIVE')),
    UNIQUE (type),
    INDEX idx_movie_type (type)
);

CREATE TABLE payment (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    bill_id BIGINT NOT NULL UNIQUE,
    amount DECIMAL(15, 2) NOT NULL,
    currency VARCHAR(10) NOT NULL DEFAULT 'VND' CHECK (currency IN ('VND', 'USD')),
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'SUCCESS', 'FAILED')),
    vnpay_transaction_id VARCHAR(255),
    response_code VARCHAR(50),
    payment_date DATETIME,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    FOREIGN KEY (bill_id) REFERENCES bill(id) ON DELETE CASCADE,
    INDEX idx_payment_bill (bill_id)
);

CREATE TABLE permission (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL UNIQUE,
    INDEX idx_permission_name (name)
);

CREATE TABLE role (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL UNIQUE,
    INDEX idx_role_name (name)
);

CREATE TABLE room_type (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL UNIQUE,
    status VARCHAR(50) NOT NULL CHECK (status IN ('ACTIVE', 'INACTIVE')),
    INDEX idx_room_type_name (name)
);

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
);

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
);

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
);

CREATE TABLE seat_type (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    UNIQUE (name),
    INDEX idx_seat_type_name (name)
);

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
);

CREATE TABLE ticket_price (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    price DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) CHECK (status IN ('ACTIVE', 'INACTIVE')),
    INDEX idx_ticket_price_price (price)
);

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
);

CREATE TABLE movie_typeofmovies (
    movie_id BIGINT NOT NULL,
    movie_type_id BIGINT NOT NULL,
    PRIMARY KEY (movie_id, movie_type_id),
    FOREIGN KEY (movie_id) REFERENCES movie(id) ON DELETE CASCADE,
    FOREIGN KEY (movie_type_id) REFERENCES movie_type(id) ON DELETE CASCADE,
    INDEX idx_movie_type_movie (movie_id),
    INDEX idx_movie_type_type (movie_type_id)
);

CREATE TABLE user_roles (
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE,
    INDEX idx_user_roles_user (user_id),
    INDEX idx_user_roles_role (role_id)
);

CREATE TABLE role_permissions (
    role_id BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,
    PRIMARY KEY (role_id, permission_id),
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permission(id) ON DELETE CASCADE,
    INDEX idx_role_permissions_role (role_id),
    INDEX idx_role_permissions_perm (permission_id)
);