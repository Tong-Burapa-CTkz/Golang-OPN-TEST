Number 2 exam 

CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE product_categories (
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    category_id INT REFERENCES categories(category_id) ON DELETE CASCADE,
    PRIMARY KEY (product_id, category_id)
);

CREATE TABLE variations (
    variation_id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    variation_type VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL
);

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    surname VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    dateofbirth DATETIME(),
    gender VARCHAR(255),
    address text,
    city VARCHAR(255),
    province VARCHAR(255),
    postno VARCHAR(255),
    tel VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE carts (
    cart_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE cart_items (
    cart_item_id SERIAL PRIMARY KEY,
    cart_id INT REFERENCES carts(cart_id) ON DELETE CASCADE,
    product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
    variation_id INT REFERENCES variations(variation_id) ON DELETE SET NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);
