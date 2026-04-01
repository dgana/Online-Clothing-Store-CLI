INSERT INTO users (email, password) VALUES
('andi@gmail.com', '123456'),
('budi@gmail.com', '123456'),
('citra@gmail.com', '123456'),
('dedi@gmail.com', '123456'),
('eka@gmail.com', '123456'),
('farah@gmail.com', '123456'),
('gina@gmail.com', '123456'),
('hadi@gmail.com', '123456'),
('intan@gmail.com', '123456'),
('joko@gmail.com', '123456');


INSERT INTO profiles (user_id, name, address, phone) VALUES
(1, 'Andi Saputra', 'Bandung', '0811111111'),
(2, 'Budi Santoso', 'Jakarta', '0822222222'),
(3, 'Citra Lestari', 'Surabaya', '0833333333'),
(4, 'Dedi Kurniawan', 'Medan', '0844444444'),
(5, 'Eka Putri', 'Yogyakarta', '0855555555'),
(6, 'Farah Aulia', 'Depok', '0866666666'),
(7, 'Gina Maharani', 'Bogor', '0877777777'),
(8, 'Hadi Pratama', 'Bekasi', '0888888888'),
(9, 'Intan Permata', 'Semarang', '0899999999'),
(10, 'Joko Widodo', 'Solo', '0810101010');

INSERT INTO profiles (user_id, name, address, phone) VALUES
(1, 'Andi Saputra', 'Bandung', '0811111111'),
(2, 'Budi Santoso', 'Jakarta', '0822222222'),
(3, 'Citra Lestari', 'Surabaya', '0833333333'),
(4, 'Dedi Kurniawan', 'Medan', '0844444444'),
(5, 'Eka Putri', 'Yogyakarta', '0855555555'),
(6, 'Farah Aulia', 'Depok', '0866666666'),
(7, 'Gina Maharani', 'Bogor', '0877777777'),
(8, 'Hadi Pratama', 'Bekasi', '0888888888'),
(9, 'Intan Permata', 'Semarang', '0899999999'),
(10, 'Joko Widodo', 'Solo', '0810101010');

INSERT INTO products (name, price, stock) VALUES
('Kaos Polos Hitam', 75000, 50),
('Kaos Polos Putih', 75000, 60),
('Hoodie Abu', 150000, 30),
('Hoodie Hitam', 155000, 25),
('Kemeja Flanel Merah', 120000, 40),
('Kemeja Flanel Biru', 120000, 35),
('Celana Jeans Slim Fit', 200000, 25),
('Celana Jeans Regular', 195000, 20),
('Jaket Denim', 250000, 15),
('Sweater Rajut', 130000, 45),
('Kaos Oversize', 90000, 55),
('Celana Chino', 180000, 30),
('Topi Baseball', 50000, 70),
('Tas Selempang', 110000, 40),
('Sepatu Sneakers', 300000, 20);

INSERT INTO orders (user_id) VALUES
(1),(2),(3),(4),(5),
(6),(7),(8),(9),(10),
(1),(2),(3),(4),(5),
(6),(7),(8),(9),(10);

INSERT INTO order_items (order_id, product_id, quantity) VALUES
(1,1,2),(1,3,1),
(2,2,1),(2,5,2),
(3,4,1),(3,6,1),
(4,7,2),(4,8,1),
(5,9,1),(5,10,2),

(6,11,1),(6,12,2),
(7,13,3),(7,1,1),
(8,2,2),(8,14,1),
(9,15,1),(9,3,1),
(10,4,2),(10,5,1),

(11,6,2),(11,7,1),
(12,8,1),(12,9,2),
(13,10,1),(13,11,2),
(14,12,1),(14,13,2),
(15,14,1),(15,15,1),

(16,1,3),(16,2,1),
(17,3,2),(17,4,1),
(18,5,1),(18,6,2),
(19,7,1),(19,8,1),
(20,9,2),(20,10,1);


===total terjual per produk===
SELECT 
    p.name,
    SUM(oi.quantity) AS total_terjual
FROM order_items oi
JOIN products p ON oi.product_id = p.id
GROUP BY p.name
ORDER BY total_terjual DESC;

=== order per customer===
SELECT 
    u.email,
    COUNT(o.id) AS total_order
FROM users u
JOIN orders o ON u.id = o.user_id
GROUP BY u.email
ORDER BY total_order DESC;


=== sisa stok per produk ===
SELECT 
    p.name,
    p.stock - IFNULL(SUM(oi.quantity),0) AS sisa_stok
FROM products p
LEFT JOIN order_items oi ON p.id = oi.product_id
GROUP BY p.id;