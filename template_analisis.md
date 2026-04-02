ERD Title: Clothing Store System
1. Entities and Their Attributes:
Entity: users

Attributes:

id (PK, AI, INT)

username (VARCHAR, UNIQUE, NOT NULL) [NEW]

email (VARCHAR, UNIQUE, NOT NULL)

password (VARCHAR, NOT NULL)

Entity: profiles

Attributes:

id (PK, AI, INT)

user_id (FK, UNIQUE, INT)

name (VARCHAR)

address (TEXT)

phone (VARCHAR)

Entity: products

Attributes:

id (PK, AI, INT)

product_code (VARCHAR, UNIQUE, NOT NULL) [NEW]

name (VARCHAR)

price (DECIMAL/FLOAT)

stock (INT)

Entity: orders

Attributes:

id (PK, AI, INT)

user_id (FK, INT)

order_date (DATETIME)

Entity: order_items

Attributes:

id (PK, AI, INT)

order_id (FK, INT)

product_id (FK, INT)

quantity (INT)

2. Relationships:
users to profiles:

Type: One to One

Description: Setiap satu akun user (diidentifikasi via username) memiliki tepat satu profile.

users to orders:

Type: One to Many

Description: Satu user dapat melakukan banyak transaksi orders, namun satu transaksi hanya dimiliki oleh satu user.

orders to order_items:

Type: One to Many

Description: Satu transaksi orders dapat memuat banyak detail barang di order_items.

products to order_items:

Type: One to Many

Description: Satu produk (diidentifikasi via product_code) dapat muncul di berbagai baris transaksi order_items.

3. Integrity Constraints:
Uniqueness: username dan email pada tabel users, serta product_code pada tabel products harus bersifat unik untuk mencegah data ganda.

Positive Value: Harga (price) pada tabel products dan quantity pada order_items harus berupa angka positif.

Referential Integrity: order_id dan product_id pada order_items harus merujuk pada ID yang valid di tabel master masing-masing.

4. Additional Notes:
Penggunaan product_code mempermudah pencarian barang di aplikasi CLI tanpa harus menghafal ID (Auto Increment).

Username digunakan sebagai identitas unik saat proses Login di aplikasi Golang.

Sistem menggunakan tabel order_items untuk memfasilitasi hubungan Many-to-Many antara pesanan dan produk pakaian.