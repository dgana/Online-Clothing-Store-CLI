package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func connectDB() {
    _ = godotenv.Load()

    dsn := os.Getenv("MYSQL_DSN")
    if dsn == "" {
        dsn = "root:password@tcp(127.0.0.1:3306)/FYGAR"
    }

    var err error
    db, err = sql.Open("mysql", dsn)

    if err != nil {
        panic(err)
    }

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Connected to MySQL!")
}

func tambahUser(email, password string) {
    query := "INSERT INTO users (email, password) VALUES (?, ?)"
    _, err := db.Exec(query, email, password)

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("User berhasil ditambahkan")
}

func tambahProduk(name string, price float64, stock int) {
    query := "INSERT INTO products (name, price, stock) VALUES (?, ?, ?)"
    _, err := db.Exec(query, name, price, stock)

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Produk berhasil ditambahkan")
}

func buatOrder(userID int) int {
    result, err := db.Exec("INSERT INTO orders (user_id) VALUES (?)", userID)

    if err != nil {
        fmt.Println("Error:", err)
        return 0
    }

    id, _ := result.LastInsertId()
    fmt.Println("Order dibuat dengan ID:", id)
    return int(id)
}

func tambahItem(orderID, productID, qty int) {
    query := "INSERT INTO order_items (order_id, product_id, quantity) VALUES (?, ?, ?)"
    _, err := db.Exec(query, orderID, productID, qty)

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Item ditambahkan ke order")
}

func laporanUser() {
    rows, _ := db.Query("SELECT id, email FROM users")

    for rows.Next() {
        var id int
        var email string
        rows.Scan(&id, &email)
        fmt.Println(id, email)
    }
}

func laporanOrder() {
    rows, _ := db.Query("SELECT id, user_id, order_date FROM orders")

    for rows.Next() {
        var id, userID int
        var date string
        rows.Scan(&id, &userID, &date)
        fmt.Println(id, userID, date)
    }
}
func laporanStok() {
    rows, _ := db.Query("SELECT id, name, stock FROM products")

    for rows.Next() {
        var id int
        var name string
        var stock int
        rows.Scan(&id, &name, &stock)
        fmt.Println(id, name, stock)
    }
}

func init(){
    connectDB()
}

func main() {
    for {
        fmt.Println("\n=== MENU ===")
        fmt.Println("1. Tambah User")
        fmt.Println("2. Tambah Produk")
        fmt.Println("3. Buat Order")
        fmt.Println("4. Tambah Item Order")
        fmt.Println("5. Laporan User")
        fmt.Println("6. Laporan Order")
        fmt.Println("7. Laporan Stok")
        fmt.Println("0. Keluar")

        var pilihan int
        fmt.Scan(&pilihan)

        switch pilihan {
        case 1:
            var email, password string
            fmt.Scan(&email, &password)
            tambahUser(email, password)

        case 2:
            var name string
            var price float64
            var stock int
            fmt.Scan(&name, &price, &stock)
            tambahProduk(name, price, stock)

        case 3:
            var userID int
            fmt.Scan(&userID)
            buatOrder(userID)

        case 4:
            var orderID, productID, qty int
            fmt.Scan(&orderID, &productID, &qty)
            tambahItem(orderID, productID, qty)

        case 5:
            laporanUser()

        case 6:
            laporanOrder()

        case 7:
            laporanStok()

        case 0:
            return
        }
    }
}