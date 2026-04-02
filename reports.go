package main

import (
	"fmt"
	"strings"
)

func laporanUser() {
	rows, err := db.Query(
		`SELECT
			u.id,
			u.username,
			IFNULL(p.name, ''),
			u.email,
			IFNULL(p.address, ''),
			IFNULL(p.phone, '')
		FROM users u
		LEFT JOIN profiles p ON p.user_id = u.id
		ORDER BY u.id`,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n=== LAPORAN USER ===")
	fmt.Printf("%-4s %-12s %-25s %-30s %-20s %-15s\n", "ID", "USERNAME", "NAME", "EMAIL", "ADDRESS", "PHONE")
	fmt.Println(strings.Repeat("-", 128))

	count := 0
	for rows.Next() {
		var id int
		var username, name, email, address, phone string
		if err := rows.Scan(&id, &username, &name, &email, &address, &phone); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%-4d %-12s %-25s %-30s %-20s %-15s\n", id, username, name, email, address, phone)
		count++
	}
	if count == 0 {
		fmt.Println("(kosong)")
	}
}

func laporanOrder() {
	rows, err := db.Query(
		`SELECT 
			o.id,
			IFNULL(p.name, ''),
			o.order_date,
			IFNULL(SUM(oi.quantity), 0) AS total_items
		FROM orders o
		LEFT JOIN profiles p ON p.user_id = o.user_id
		LEFT JOIN order_items oi ON oi.order_id = o.id
		GROUP BY o.id, p.name, o.order_date
		ORDER BY o.id`,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n=== LAPORAN ORDER ===")
	fmt.Printf("%-6s %-25s %-20s %-10s\n", "ORDER", "CUSTOMER", "DATE", "ITEMS")
	fmt.Println(strings.Repeat("-", 66))

	count := 0
	for rows.Next() {
		var id int
		var name, date string
		var totalItems int
		if err := rows.Scan(&id, &name, &date, &totalItems); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%-6d %-25s %-20s %-10d\n", id, name, date, totalItems)
		count++
	}
	if count == 0 {
		fmt.Println("(kosong)")
	}
}

func laporanStok() {
	rows, err := db.Query(
		`SELECT
			p.id,
			p.product_code,
			p.name,
			p.price,
			p.stock,
			IFNULL(SUM(oi.quantity), 0) AS sold,
			p.stock - IFNULL(SUM(oi.quantity), 0) AS remaining
		FROM products p
		LEFT JOIN order_items oi ON oi.product_id = p.id
		GROUP BY p.id, p.product_code, p.name, p.price, p.stock
		ORDER BY p.id`,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n=== LAPORAN STOK ===")
	fmt.Printf("%-4s %-8s %-30s %-12s %-8s %-8s %-10s\n", "ID", "CODE", "PRODUCT", "PRICE", "STOCK", "SOLD", "REMAIN")
	fmt.Println(strings.Repeat("-", 92))

	count := 0
	for rows.Next() {
		var id int
		var code, name string
		var price float64
		var stock, sold, remaining int
		if err := rows.Scan(&id, &code, &name, &price, &stock, &sold, &remaining); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%-4d %-8s %-30s %-12.2f %-8d %-8d %-10d\n", id, code, name, price, stock, sold, remaining)
		count++
	}
	if count == 0 {
		fmt.Println("(kosong)")
	}
}
