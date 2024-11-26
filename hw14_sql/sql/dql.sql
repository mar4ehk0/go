-- Напишите запрос на выборку пользователей и выборку товаров
SELECT u.name, p.name, p.price FROM users u
JOIN orders o ON u.id = o.user_id 
JOIN orders_products op ON o.id = op.order_id 
JOIN products p ON p.id = op.product_id 

-- Напишите запрос на выборку статистики по пользователю (общая сумма заказов/средняя цена товара)
SELECT o.user_id, SUM(o.total_amount) AS total FROM orders o WHERE o.user_id = 2 GROUP BY o.user_id 

SELECT AVG(p.price) AS average_price FROM products p


SELECT * FROM orders WHERE total_amount=200;

SELECT id, name, price FROM products WHERE price=200;

SELECT id, name, email, password FROM users WHERE email='lorem4@example.com';