CREATE TABLE IF NOT EXISTS products (
    sku VARCHAR PRIMARY KEY,
    description VARCHAR(255),
    image TEXT,
    price NUMERIC
);

CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    customer INT NOT NULL,
    product VARCHAR NOT NULL,
    quantity INT DEFAULT 1,
    CONSTRAINT fk_customer_id FOREIGN KEY (customer) REFERENCES customers (id),
    CONSTRAINT fk_product_id FOREIGN KEY (product) REFERENCES products (sku)
);
INSERT INTO products(sku, description, image, price)
VALUES ('218388', 'Fone De Ouvido Branco JBLT115BTW', 'https://d3ddx6b2p2pevg.cloudfront.net/Custom/Content/Products/11/01/1101082_fone-de-ouvido-jblt115btw-branco_z1_637794143502654247.jpg', 169.0);


INSERT INTO customers(id, name)
VALUES (1, 'Erick');

INSERT INTO orders(id, customer, product, quantity)
VALUES (1, 1, '218388', 2);


-- CREATE TABLE users (
--   id serial,
--   username VARCHAR(25) NOT NULL,
--   enabled boolean DEFAULT TRUE,
--   last_login timestamp NOT NULL DEFAULT NOW(),
--   PRIMARY KEY (id)
-- );

-- /*
--  one to one: User has one address
-- */

-- CREATE TABLE addresses (
--   user_id int NOT NULL,
--   street VARCHAR(30) NOT NULL,
--   city VARCHAR(30) NOT NULL,
--   state VARCHAR(30) NOT NULL,
--   PRIMARY KEY (user_id),
--   CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
-- );
