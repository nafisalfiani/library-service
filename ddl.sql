CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL UNIQUE,
    full_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    deposit_amount DECIMAL DEFAULT 0
);

CREATE TABLE deposit_histories (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    amount DECIMAL,
    type VARCHAR(255),
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    stock_availability INTEGER NOT NULL,
    rental_cost DECIMAL NOT NULL,
    category_id INTEGER REFERENCES categories(id),
    UNIQUE (name, category_id)
);

CREATE TABLE rentals (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    rental_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE rental_details (
    id SERIAL PRIMARY KEY,
    rental_id INTEGER REFERENCES rentals(id),
    book_id INTEGER REFERENCES books(id),
    return_date TIMESTAMP,
    returned BOOLEAN DEFAULT FALSE,
    UNIQUE (rental_id, book_id)
);

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    rental_id INTEGER REFERENCES rentals(id),
    amount DECIMAL NOT NULL,
    payment_method VARCHAR(255) NOT NULL,
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
