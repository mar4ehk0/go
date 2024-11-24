-- products
INSERT INTO products ("name", "price")
VALUES
('products1',  100),
('products2',  200),
('products3',  300),
('products4',  400);

UPDATE products SET name='products200' WHERE price=200;

DELETE FROM products WHERE name='products200';

-- orders
INSERT INTO orders ("user_id", "order_date", "total_amount")
VALUES
(2, '1999-01-08 04:05:06 -8:00', 100),
(2, '1999-01-08 04:06:06 -8:00', 200),
(3, '1999-01-08 05:05:06 -8:00', 100),
(4, '1999-01-08 05:06:06 -8:00', 500);

UPDATE orders SET total_amount=600 WHERE id=2;

DELETE FROM orders WHERE id=2;

-- orders_products
INSERT INTO orders_products (order_id, product_id)
VALUES
(2, 2),
(2, 3),
(2, 4);

-- users
INSERT INTO users ("name", "email", "password")
VALUES
('lorem ipsum1', 'lorem1@example.com', 'password'),
('lorem ipsum2', 'lorem2@example.com', 'password'),
('lorem ipsum3', 'lorem3@example.com', 'password'),
('lorem ipsum4', 'lorem4@example.com', 'password');

UPDATE users SET password='new_password' WHERE email='lorem2@example.com';

DELETE FROM users WHERE name='lorem ipsum1';